package util

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate(c *fiber.Ctx, out interface{}) error {
	// Parse
	if err := c.BodyParser(out); err != nil {
		return err
	}

	// Validate
	if err := validator.New().Struct(out); err != nil {
		validationError := err.(validator.ValidationErrors)[0]
		return errors.New(validationError.Field() + ": " + validationError.Tag())
	}

	return nil
}
