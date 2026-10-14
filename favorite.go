package qobuzsdk

type Favorite struct {
	TrackIds  []int    `schema:"track_ids"`
	AlbumIds  []string `schema:"album_ids"`
	ArtistIds []int    `schema:"artist_ids"`
}
