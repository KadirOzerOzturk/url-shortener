package entities

type ShortenRequest struct {
	OriginalUrl string `json:"original_url" validate:"required"`
}
