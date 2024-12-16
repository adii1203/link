package types

type Link struct {
	Id             int
	DestinationUrl string `json:"destination_url" validate:"required,url"`
}
