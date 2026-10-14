package qobuzsdk

type Client interface {
	LoginWithOAuth(oAuthRedirectUrl string) OAuthLogin
	LoginWithToken(token string)
	IsAuthenticated() bool

	GetTrack(trackId int) (Track, error)
	GetTrackInto(trackId int, mapTo any) error
	SearchTracks(query string, limit, offset int16) (TracksSearchResult, error)
	SearchTracksInto(query string, limit, offset int16, mapTo any) error
	StreamTrack(trackId int, format Format) (TrackFileUrl, error)
	StreamTrackInto(trackId int, format Format, mapTo any) error
	AddTrackToFavorites(trackId int) error
	RemoveTrackFromFavorites(trackId int) error

	GetAlbum(albumId string) (Album, error)
	GetAlbumInto(albumId string, mapTo any) error
	SearchAlbums(query string, limit, offset int16) (AlbumsSearchResult, error)
	SearchAlbumsInto(query string, limit, offset int16, mapTo any) error
	AddAlbumToFavorites(albumId string) error
	RemoveAlbumFromFavorites(albumId string) error

	GetPlaylist(playlistId int) (Playlist, error)
	GetPlaylistInto(playlistId int, mapTo any) error

	SearchPlaylists(query string, limit, offset int16) (PlaylistsSearchResult, error)
	SearchPlaylistsInto(query string, limit, offset int16, mapTo any) error

	GetArtist(artistId int) (ArtistPage, error)
	GetArtistInto(artistId int, mapTo any) error
	SearchArtists(query string, limit, offset int16) (ArtistsSearchResult, error)
	SearchArtistsInto(query string, limit, offset int16, mapTo any) error
	AddArtistToFavorites(artistId int) error
	RemoveArtistFromFavorites(artistId int) error

	AddToFavorites(favorite Favorite) error
	RemoveFromFavorites(favorite Favorite) error

	SearchCatalog(query string, limit, offset int16) (CatalogSearchResult, error)
	SearchCatalogInto(query string, limit, offset int16, mapTo any) error
}
