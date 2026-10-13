package qobuzsdk

type Pagination[T any] struct {
	Limit  int16 `json:"limit"`
	Offset int16 `json:"offset"`
	Total  int32 `json:"total"`
	Items  []T   `json:"items"`
}
