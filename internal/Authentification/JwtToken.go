package Authentification

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v3"
	"os"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email, secretKey string, expTime time.Time) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return fmt.Errorf("token is expired")
		}
		return nil
	}

	return fmt.Errorf("invalid token")
}

func JWTAuthMiddleware(c fiber.Ctx) error {
	cookie := c.Cookies("jwt-dnz", "def")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	err := ValidateToken(cookie, os.Getenv("JWT_KEY"))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Token is valid
	return c.Next()
}

func TokenExpired(err error) bool {
	if err.Error() == "token is expired" {
		return true
	}
	return false
}

func InvalidToken(err error) bool {
	if err.Error() == "invalid token" {
		return true
	}
	return false
}
