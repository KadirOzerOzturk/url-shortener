package entities

type ShortenRequest struct {
	OriginalUrl string `json:"original_url" validate:"required"`
	UserEmail   string `json:"user_email" validate:"required,email"`
}
