package entities

import "time"

type Url struct {
	Base
	OriginalUrl  string    `json:"original_url" validate:"required"`
	ShortenedUrl string    `json:"shortened_url" validate:"required"`
	UsageCount   uint      `json:"usage_count" validate:"required"`
	ExpiresAt    time.Time `json:"expires_at" validate:"required"`
}
