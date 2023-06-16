package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}