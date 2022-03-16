package route

import (
	controller "github.com/beenson/URL_Shortener/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	// route /api/v1
	route := app.Group("/api/v1")
	route.Post("/urls", controller.CreateShortenURL)

	// Redirect
	app.Get("/:url_id", controller.Redirect)
}
