package handlers

import (
	"bytes"
	"context"
	lx "github.com/Nevoral/LuxeGo"
	"github.com/gofiber/fiber/v3"
	"log"
	"path/filepath"
)

const pathBase = "./web"

func SendJs(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, "static", "js", c.Params("*"))
	return sendFile(c, filePath)
}

func SendAsset(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, "static", "assets", c.Params("*"))
	return sendFile(c, filePath)
}

func SendCss(c fiber.Ctx) error {
	filePath := filepath.Join(pathBase, "static", "css", c.Params("*"))
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
		// Set the HTTP Content-Type header.
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

		// Buffer to store the rendered HTML.
		var output bytes.Buffer
		for _, value := range page {
			// Render the page into the buffer.
			if err := value.Render(context.Background(), &output); err != nil {
				// Handle the rendering error, e.g., by logging and returning an HTTP 500 error.
				log.Printf("Error rendering page: %v", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}

		// Send the rendered HTML as the response.
		return c.SendStream(&output)
	}
}
