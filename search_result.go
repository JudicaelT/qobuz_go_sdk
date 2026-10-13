package qobuzsdk

type CatalogSearchResult struct {
	Query     string                       `json:"query"`
	Tracks    Pagination[Track]            `json:"tracks"`
	Albums    Pagination[Album]            `json:"albums"`
	Playlists Pagination[Playlist]         `json:"playlists"`
	Artists   Pagination[ArtistSearchItem] `json:"artists"`
}

type TracksSearchResult struct {
	Query  string            `json:"query"`
	Tracks Pagination[Track] `json:"tracks"`
}

type AlbumsSearchResult struct {
	Query  string            `json:"query"`
	Albums Pagination[Album] `json:"albums"`
}

type PlaylistsSearchResult struct {
	Query     string               `json:"query"`
	Playlists Pagination[Playlist] `json:"playlists"`
}

type ArtistsSearchResult struct {
	Query   string                       `json:"query"`
	Artists Pagination[ArtistSearchItem] `json:"artists"`
}

type ArtistSearchItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Picture     string `json:"picture"`
	AlbumsCount int    `json:"albums_count"`
	Image       Image  `json:"image"`
}
