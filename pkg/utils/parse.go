package util

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidate(c *fiber.Ctx, out interface{}) error {
	// Parse
	if err := c.BodyParser(out); err != nil {
		return err
	}

	// Create validator
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	// Validate
	if err := validate.Struct(out); err != nil {
		validationError := err.(validator.ValidationErrors)[0]
		conj := ""
		if validationError.Tag() == "required" {
			conj = " is "
		} else {
			conj = " should be "
		}
		return errors.New(validationError.Field() + conj + validationError.Tag())
	}

	return nil
}
