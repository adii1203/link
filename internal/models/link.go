package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	DestinationUrl string   `json:"destination_url"`
	Slug           string   `json:"slug" gorm:"uniqueIndex"`
	Metadata       Metadata `json:"metadata" gorm:"embedded"`
}
type Metadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
