package types

type Link struct {
	Id             int
	DestinationUrl string `json:"destination_url" validate:"required,url"`
	Slug           string `json:"slug" validate:"ascii"`
}
