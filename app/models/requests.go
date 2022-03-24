package model

type ShortenUrlResquest struct {
	URL      string `example:"<original_url>" json:"url" validate:"required,url"`
	ExpireAt string `example:"2021-02-08T09:20:41Z" json:"expireAt" validate:"required"`
}
