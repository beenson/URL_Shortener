package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate(c *fiber.Ctx, out interface{}) error {
	// Parse
	if err := c.BodyParser(out); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Validate
	if err := validator.New().Struct(out); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil
}
