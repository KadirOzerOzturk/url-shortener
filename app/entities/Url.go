package entities

import "time"

type Url struct {
	Base
	OriginalUrl  string    `json:"original_url" `
	ShortenedUrl string    `json:"shortened_url" `
	UsageCount   uint      `json:"usage_count" `
	ExpiresAt    time.Time `json:"expires_at" `
}
