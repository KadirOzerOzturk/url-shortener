package entities

import "time"

type Log struct {
	ShortenedUrl string    `json:"shortened_url "gorm:"primary_key `
	AccessedIp   string    `json:"accessed_ip" `
	AccessedAt   time.Time `json:"accessed_at" `
	AccessCount  int       `json:"access_count" `
}
