package Authentification

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/mail"
	"unicode"
)

func UsernameTest(Username string) (int, error) {
	if len(Username) < 3 {
		return fiber.StatusNotAcceptable, fmt.Errorf("uživatelské jméno je příliš krátké")
	}
	return fiber.StatusOK, nil
}

func EmailTest(email string) (int, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return fiber.StatusNotAcceptable, err
	}
	return fiber.StatusOK, nil
}

func PasswordTest(password string) (int, error) {
	msg := "heslo neobsahuje: "
	cond := map[string]bool{
		"min 1 A-Z":        false,
		"min 1 @#$%^&*...": false,
		"min 1 a-z":        false,
		"min 1 0-9":        false,
		"min 8 charakterů": false,
	}
	if len(password) >= 8 {
		cond["min 8 charakterů"] = true
	}

	for _, r := range password {
		if unicode.IsUpper(r) {
			cond["min 1 A-Z"] = true
			continue
		} else if unicode.IsLower(r) {
			cond["min 1 a-z"] = true
			continue
		} else if unicode.IsDigit(r) {
			cond["min 1 0-9"] = true
			continue
		}
		cond["min 1 @#$%^&*..."] = true
	}
	var missing []string
	for key, val := range cond {
		if !val {
			missing = append(missing, key)
			continue
		}
	}
	for i, val := range missing {
		if i+1 == len(missing) {
			msg += val + "."
			break
		}
		if i+2 == len(missing) {
			msg += val + " a "
			continue
		}
		msg += val + ", "
	}
	fmt.Println(msg)
	if len(missing) > 0 {
		return fiber.StatusNotAcceptable, fmt.Errorf(msg)
	}
	return fiber.StatusOK, nil
}
