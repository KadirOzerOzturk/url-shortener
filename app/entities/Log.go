package entities

import "time"

type Log struct {
	ShortenedUrl string    `json:"shortened_url ,"gorm:"primary_key, validate:"required"`
	AccessedIp   string    `json:"accessed_ip" validate:"required"`
	AccessedAt   time.Time `json:"accessed_at" validate:"required"`
	AccessCount  int       `json:"access_count" default0:"1" validate:"required"`
}
