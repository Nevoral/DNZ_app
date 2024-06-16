package server

import (
	"github.com/Nevoral/DNZ_app/internal/database"
	"github.com/gofiber/fiber/v3"
)

type FiberServer struct {
	*fiber.App
	db database.DbClient
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
		db:  database.New(),
	}

	return server
}
