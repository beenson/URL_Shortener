package controller

import (
	"time"

	model "github.com/beenson/URL_Shortener/app/models"
	"github.com/beenson/URL_Shortener/pkg/repository"
	util "github.com/beenson/URL_Shortener/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Create Shorten URL
// @Description  generate a shorten URL
// @Tags         url
// @Accept       json
// @Produce      json
// @Param        body  body  model.ShortenUrlResquest  true  "Shorten URL information"
// @Success      200  {object}  model.ShortenUrlResponse  "success"
// @Failure      400  {object}  model.HTTPError           "wrong type or missing value"
// @Failure      500  {object}  model.HTTPError           "server error"
// @Router       /api/v1/urls [POST]
func CreateShortenURL(c *fiber.Ctx) error {
	// Parse
	shortenURL := &model.ShortenUrlResquest{}
	if err := util.ParseAndValidate(c, shortenURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&model.HTTPError{
			Message: err.Error(),
		})
	}

	// Parse Time
	expire_at, err := time.Parse(time.RFC3339, shortenURL.ExpireAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&model.HTTPError{
			Message: "please check expireAt format",
		})
	}

	// Check if expire time has passed
	if expire_at.Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(&model.HTTPError{
			Message: "expire time has passed",
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
				return c.Status(fiber.StatusInternalServerError).JSON(&model.HTTPError{
					Message: err.Error(),
				})
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

// @Summary      Redirect to URL
// @Description  redirect to origin url if {url_id} exist and without expired
// @Tags         url
// @Param        url_id  path  string  true  "The id which response by /api/v1/urls"
// @Success      302  "redirect"
// @Failure      404  "{url_id} not found"
// @Router       /{url_id} [GET]
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
