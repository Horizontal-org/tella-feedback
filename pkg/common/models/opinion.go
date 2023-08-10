package models

import "gorm.io/gorm"

type Opinion struct {
	gorm.Model
	Text string `json:"text"`
	Platform string `json:"platform"`
}