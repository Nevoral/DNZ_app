package main

import (
	"fmt"
	zlog "github.com/Nevoral/DNZ_app/internal/logging"
	"github.com/Nevoral/DNZ_app/internal/server"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

//
//type Product struct {
//	ID    int     `json:"id"`
//	Title string  `json:"title"`
//	Price float64 `json:"price"`
//	Count int     `json:"count"`
//}
//
//var products []Product
//
//func saveHandler(c *fiber.Ctx) error {
//	file, err := json.MarshalIndent(products, "", " ")
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to marshal products")
//	}
//	err = os.WriteFile("state.json", file, 0644)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
//	}
//	return c.SendString("State saved")
//}
//
//func loadHandler(c *fiber.Ctx) error {
//	file, err := os.ReadFile("state.json")
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read file")
//	}
//	err = json.Unmarshal(file, &products)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to unmarshal products")
//	}
//	return c.SendString("State loaded")
//}
//
//func main() {
//	app := fiber.New()
//
//	// Serve all files in the "public" directory
//	app.Static("/", "./public")
//
//	// GET request to root will serve the index.html
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendFile("./public/index.html")
//	})
//
//	app.Post("/save", saveHandler)
//	app.Get("/load", loadHandler)
//
//	app.Post("/log", func(c *fiber.Ctx) error {
//		// Increment and log to file
//		//fmt.Println(string(c.Body()))
//		logToFile(string(c.Body()) + "\n")
//
//		return c.SendString("OK")
//	})
//
//	app.Listen(":3000")
//}
//
//func logToFile(message string) {
//	file, err := os.OpenFile("orders.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//
//	if _, err := file.WriteString(message); err != nil {
//		fmt.Println(err)
//	}
//}

func main() {
	zlog.InitLogger()
	if err := godotenv.Load("./.env"); err != nil {
		zlog.FatalLog(fmt.Sprintf("loading .env file: %v", err))
	}

	s := server.New()

	s.Router()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := s.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
