package api

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber/api/middleware"
	"go_fiber/internal/app/handler"
)

// SetupRoutes initializes and configures all API routes
func SetupRoutes(app *fiber.App) {
	// Group routes if needed
	api := app.Group("/api")

	// Apply middleware to the entire group
	api.Use(middleware.AuthMiddleware)

	// Define individual routes
	api.Get("/task", handler.GetTasks)
	api.Get("/task/:id", handler.GetTasksByID)
	api.Post("/task", handler.CreateTask)
	api.Put("/task", handler.UpdateTask)
	api.Delete("/task/:id", handler.DeleteTask)
}
