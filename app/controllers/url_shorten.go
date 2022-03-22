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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Parse Time
	expire_at, err := time.Parse(time.RFC3339, shortenURL.ExpireAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "please check expireAt format",
		})
	}

	// Check if expire time has passed
	if expire_at.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "expire time has passed",
		})
	}

	// Create Shorten
	shorten := &model.Shorten{
		URL:      shortenURL.URL,
		ExpireAt: expire_at,
	}

	// Generate Code and Try To Create
	codeAvaliable := false
	for len := repository.Default_code_length; !codeAvaliable; len++ {
		for count := 0; count < repository.Maximum_tries; count++ {
			shorten.Code = util.GenerateCode(len)
			if err := model.CreateShorten(shorten); err == repository.ErrCodeUnavailable {
				continue
			} else if err != nil {
				return c.SendStatus(fiber.StatusBadRequest)
			} else {
				codeAvaliable = true
				break
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(&model.ShortenUrlResponse{
		ID:       shorten.Code,
		ShortUrl: repository.Host_address + "/" + shorten.Code,
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
