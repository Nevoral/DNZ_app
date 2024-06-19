package handlers

import (
	"bytes"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/gofiber/fiber/v3"
	"log"
	"path/filepath"
)

const pathBase = "./web"

func SendJs(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, c.Query("path", ""), "static", "js", c.Query("name"))
	return sendFile(c, filePath)
}

func SendAsset(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, c.Query("path", ""), "static", "assets", c.Query("name"))
	return sendFile(c, filePath)
}

func SendCss(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, c.Query("path", ""), "static", "css", c.Query("name"))
	return sendFile(c, filePath)
}

func sendFile(c fiber.Ctx, path string) error {
	err := c.SendFile(path)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.SendString("File not found")
	}
	return nil
}

func SendHTML(page ...lx.Content) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

		var output bytes.Buffer
		for _, value := range page {
			if err := value.Render(c.UserContext(), &output); err != nil {
				log.Printf("Error rendering page: %v", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}
		return c.SendStream(&output)
	}
}
