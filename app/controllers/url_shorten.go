package controller

import (
	"time"

	model "github.com/beenson/URL_Shortener/app/models"
	"github.com/beenson/URL_Shortener/pkg/repository"
	util "github.com/beenson/URL_Shortener/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateShortenURL(c *fiber.Ctx) error {
	// Parse
	shortenURL := &model.ShortenUrlResquest{}
	if err := util.ParseAndValidate(c, shortenURL); err != nil {
		return err
	}

	// Parse Time
	t, err := time.Parse(time.RFC3339, shortenURL.ExpireAtString)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Create Shorten
	shorten := &model.Shorten{
		Code:     util.GenerateCode(repository.Default_code_length),
		URL:      shortenURL.URL,
		ExpireAt: t,
	}
	if a := model.CreateShorten(shorten); a != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Id":       shorten.Code,
		"shortUrl": repository.Host_address + "/" + shorten.Code,
	})
}

func Redirect(c *fiber.Ctx) error {
	shorten := &model.Shorten{
		Code: c.Params("url_id"),
	}

	if shorten.Code == "favicon.ico" {
		return nil
	}

	if err := model.GetOriginalUrl(shorten); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Redirect(shorten.URL)
}
