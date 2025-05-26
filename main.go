package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/koriebruh/gok8s-deploy/config"
)

func main() {
	// Initialize Fiber app with custom configuration
	config.LoadConfig()

	app := fiber.New(
		config.FiberConfig,
	)
	app.Use(config.RecoverMiddleware) //Handle panic recovery
	app.Use(config.LoggerConfig)      // Log requests
	app.Use(config.SecurityConfig)    // Security headers
	app.Use(config.CompressionConfig) // Compress responses for more fast delivery
	app.Use(config.CORSConfig)        // Enable CORS for all routes
	app.Use(config.RateLimitConfig)   // Rate limiting to prevent abuse
	app.Use(etag.New())               // ETag middleware for caching

	// Registration of API routes
	api := app.Group("/api/v1")
	api.Use(config.APIKeyMiddleware())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "message": "Welcome to the API!"})
	})

	// Tambahkan rute untuk /health dan /metrics jika perlu
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"status":  "healthy",
		})
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
