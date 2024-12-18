package app

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber/api"
	"go_fiber/api/middleware"
	"go_fiber/database"
	"go_fiber/internal/config"
)

func Initialize() *fiber.App {
	app := fiber.New()

	config.Load()

	database.Migrate()

	app.Use(middleware.LoggingMiddleware)

	api.SetupRoutes(app)

	return app
}
