package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	DestinationUrl string
	Key            string
}
