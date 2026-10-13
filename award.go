package qobuzsdk

type Award struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	AwardSlug       string `json:"award_slug"`
	AwardedAt       int    `json:"awarded_at"`
	AwardId         string `json:"award_id"`
	PublicationId   string `json:"publication_id"`
	PublicationName string `json:"publication_name"`
	PublicationSlug string `json:"publication_slug"`
}
