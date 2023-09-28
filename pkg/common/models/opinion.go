package models

import "gorm.io/gorm"

// @Description Opinion model
type Opinion struct {
	gorm.Model
	Text string `json:"text"`
	Platform string `json:"platform"`
}