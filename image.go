package qobuzsdk

type Image struct {
	Small      string `json:"small"`
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extralarge"`
	Mega       string `json:"mega"`
}

type CoverImage struct {
	Small     string `json:"small"`
	Thumbnail string `json:"thumbnail"`
	Large     string `json:"large"`
	Back      string `json:"back"`
}
