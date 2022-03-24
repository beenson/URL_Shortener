package model

type ShortenUrlResponse struct {
	ID       string `example:"<url_id>" json:"id"`
	ShortUrl string `example:"http://localhost/<url_id>" json:"shortUrl"`
}

type HTTPError struct {
	Message string `example:"error message" json:"message"`
}
