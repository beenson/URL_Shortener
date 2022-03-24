package route

import (
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/beenson/URL_Shortener/docs"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App) {
	// Setup host address
	docs.SwaggerInfo.Host = os.Getenv("HOST_ADDRESS")

	// Swagger init
	app.Get("/swagger/*", swagger.HandlerDefault)
}
