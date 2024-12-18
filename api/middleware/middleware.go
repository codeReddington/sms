package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	// Record the start time of the request
	start := time.Now()

	// Continue processing the request
	err := c.Next()

	// Calculate the request duration
	duration := time.Since(start)

	// Log the request details
	log.Printf(
		"Request: %s %s | Status: %d | Duration: %v",
		c.Method(),
		c.Path(),
		c.Response().StatusCode(),
		duration,
	)

	// Log errors if they occurred
	if err != nil {
		log.Printf("Error: %v", err)
	}

	return err
}

// AuthMiddleware is a middleware function for authentication and authorization
func AuthMiddleware(c *fiber.Ctx) error {
	// Example: Check if the user is authenticated (e.g., by verifying a JWT token)
	if !isAuthenticated(c) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Example: Check if the authenticated user has the necessary permissions
	if !hasPermission(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	// Attach user information to the request context if needed
	c.Locals("user", getAuthenticatedUser(c))

	// Continue processing the request
	return c.Next()
}

// Helper functions for authentication and authorization checks

func isAuthenticated(c *fiber.Ctx) bool {
	// Implement your authentication logic here (e.g., check JWT token)
	// Return true if the user is authenticated, false otherwise
	// Example: return checkJWTToken(c)
	return true
}

func hasPermission(c *fiber.Ctx) bool {
	// Implement your authorization logic here (e.g., check user roles)
	// Return true if the user has permission, false otherwise
	// Example: return checkUserRole(c, "admin")
	return true
}

func getAuthenticatedUser(c *fiber.Ctx) interface{} {
	// Implement a function to fetch user information from the authentication
	// context (e.g., JWT claims or user session) and return it
	// Example: return getUserFromJWT(c)
	return nil
}
