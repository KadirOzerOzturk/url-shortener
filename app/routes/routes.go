package routes

import (
	"github.com/KadirOzerOzturk/url-shortener/app/controllers/urls"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	group := app.Group("/url")
	{

		group.Post("/shorten", urls.Shorten)
		group.Get("/", urls.Index)
		group.Get("/:short_url", urls.Redirect)

		group.Get("/:short_url/stats", urls.UrlStats)
		group.Delete("/:short_url", urls.Delete)
	}
}
