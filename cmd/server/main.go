package main

import (
	"time"

	"github.com/KadirOzerOzturk/url-shortener/app/routes"
	_ "github.com/KadirOzerOzturk/url-shortener/internal/database"
	_ "github.com/KadirOzerOzturk/url-shortener/internal/migrations"
	"github.com/KadirOzerOzturk/url-shortener/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Fiber'ın kendi cors middleware'ini import ediyoruz
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/joho/godotenv/autoload"
)

var config = fiber.Config{
	BodyLimit:    1024 * 1024 * 1024,
	ErrorHandler: server.ErrorHandler,
}

func main() {
	// Fiber uygulaması başlat
	app := fiber.New(config)

	// CORS middleware'ini ekle
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://shorterly.net/,https://url-shortener-znmr.onrender.com",
		AllowMethods: "GET,POST, PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Rate limiter middleware'i
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString("Rate limit exceeded. Please try again later.")
		},
	}))

	// Route'ları ayarla
	routes.SetupRoutes(app)

	// Uygulamayı başlat
	app.Listen(":8080")
}
