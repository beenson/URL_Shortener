package model

type ShortenUrlResquest struct {
	URL      string `json:"url" validate:"required,url"`
	ExpireAt string `json:"expireAt" validate:"required"`
}
