package qobuzsdk

type Artist struct {
	Id                           int               `json:"id"`
	Name                         string            `json:"name"`
	AlbumsAsPrimaryArtistCount   int               `json:"albums_as_primary_artist_count"`
	AlbumsAsPrimaryComposerCount int               `json:"albums_as_primary_composer_count"`
	AlbumsCount                  int               `json:"albums_count"`
	Slug                         string            `json:"slug"`
	Picture                      string            `json:"picture"`
	Image                        CoverImage        `json:"image"`
	SimilarArtistIds             []int             `json:"similar_artist_ids"`
	Biography                    ArtistBiography   `json:"biography"`
	AlbumLastRelease             Album             `json:"album_last_release"`
	Albums                       Pagination[Album] `json:"albums"`
}

type ArtistBiography struct {
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Source   string `json:"source"`
	Language string `json:"language"`
}

type ArtistPage struct {
	// TODO add images, similar_artists, last_release, releases, playlists
	Id              int                   `json:"id"`
	Name            ArtistPageDisplayName `json:"name"`
	ArtistCategory  string                `json:"artist_category"`
	Biography       ArtistPageBiography   `json:"biography"`
	TopTracks       []Track               `json:"top_tracks"`
	TracksAppearsOn []Track               `json:"tracks_appears_on"`
}

type ArtistPageDisplayName struct {
	Display string `json:"display"`
}

type ArtistPageBiography struct {
	Content  string `json:"content"`
	Source   string `json:"source"`
	Language string `json:"language"`
}
