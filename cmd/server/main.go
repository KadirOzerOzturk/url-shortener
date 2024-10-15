package main

import (
	"time"

	"github.com/KadirOzerOzturk/url-shortener/app/routes"

	_ "github.com/KadirOzerOzturk/url-shortener/internal/database"
	_ "github.com/KadirOzerOzturk/url-shortener/internal/migrations"
	"github.com/KadirOzerOzturk/url-shortener/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	redisStorage "github.com/gofiber/storage/redis"
	_ "github.com/joho/godotenv/autoload"
)

var config = fiber.Config{

	BodyLimit: 1024 * 1024 * 1024,

	ErrorHandler: server.ErrorHandler,
}

func main() {
	app := fiber.New(config)
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
		Storage: redisStorage.New(redisStorage.Config{
			Host:     "localhost",
			Port:     6379,
			Password: "",
		}),
	}))

	routes.SetupRoutes(app)
	app.Listen(":3000")
}
