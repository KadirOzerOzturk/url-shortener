package routes

import (
	"github.com/KadirOzerOzturk/url-shortener/app/controllers/auths"
	"github.com/KadirOzerOzturk/url-shortener/app/controllers/mails"
	"github.com/KadirOzerOzturk/url-shortener/app/controllers/urls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	// Enable CORS
	app.Use(cors.New())

	group := app.Group("/url")
	{
		group.Post("/shorten", urls.Shorten)
		group.Get("/", urls.Index)
		group.Get("/:short_url", urls.Redirect)
		group.Get("/:short_url/stats", urls.UrlStats)
		group.Delete("/:short_url", urls.Delete)
		group.Get("/get_by_email/:email", urls.GetUrlsByUser)
	}
	mail := app.Group("/mail")
	{
		mail.Post("/send", mails.SendMail)
	}
	auth := app.Group("/auth")
	{
		auth.Post("/login", auths.Login)
		auth.Post("/register", auths.Register)
	}
}
