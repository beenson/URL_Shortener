package route

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(
		// Anonymous function
		func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusNotFound)
		},
	)
}
