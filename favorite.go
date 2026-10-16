package qobuzsdk

type Favorite struct {
	TrackIds  []int    `form:"track_ids"`
	AlbumIds  []string `form:"album_ids"`
	ArtistIds []int    `form:"artist_ids"`
}
