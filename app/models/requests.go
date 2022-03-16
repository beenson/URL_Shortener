package model

type ShortenUrlResquest struct {
	URL            string `json:"url" validate:"required,url"`
	ExpireAtString string `json:"expireAt" validate:"required"`
}
