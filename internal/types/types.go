package types

import "github.com/adii1203/link/internal/models"

type Link struct {
	Id             uint            `json:"id"`
	DestinationUrl string          `json:"destination_url" validate:"required,url"`
	Slug           string          `json:"slug" validate:"ascii"`
	Metadata       models.Metadata `json:"metadata"`
}
