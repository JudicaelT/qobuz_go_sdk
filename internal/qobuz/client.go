package qobuz

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/rest"
)

const (
	BaseUrl     = "https://www.qobuz.com"
	ApiBaseUrl  = "https://www.qobuz.com/api.json/0.2"
	TokenHeader = "X-User-Auth-Token"
)

type Client struct {
	apiClient           rest.ApiClient
	appId               int32
	appSecret           string
	privateKey          string
	getCurrentTimestamp func() int64
}

func NewClient(
	apiClient rest.ApiClient,
	appId int32,
	appSecret string,
	privateKey string,
	getCurrentTimestamp func() int64,
) *Client {
	return &Client{
		apiClient:           apiClient,
		appId:               appId,
		appSecret:           appSecret,
		privateKey:          privateKey,
		getCurrentTimestamp: getCurrentTimestamp,
	}
}

func (self *Client) LoginWithOAuth(oAuthRedirectUrl string) qobuzsdk.OAuthLogin {
	return qobuzsdk.OAuthLogin{
		LoginLink: fmt.Sprintf(
			"%s/signin/oauth?ext_app_id=%d&redirect_url=%s",
			BaseUrl,
			self.appId,
			oAuthRedirectUrl,
		),
		OnCodeReceived: func(code string) error {
			endpoint := fmt.Sprintf(
				"oauth/callback?code=%s&private_key=%s",
				code,
				self.privateKey,
			)
			var exchange TokenExchange
			if err := self.apiClient.Get(endpoint, &exchange); err != nil {
				return err
			}
			self.LoginWithToken(exchange.Token)
			return nil
		},
	}
}

func (self *Client) LoginWithToken(token string) {
	self.apiClient.SetHeader(TokenHeader, token)
}

func (self *Client) IsAuthenticated() bool {
	return self.apiClient.HasHeader(TokenHeader)
}

func (self *Client) GetTrack(trackId int) (qobuzsdk.Track, error) {
	var track qobuzsdk.Track
	err := self.GetTrackInto(trackId, &track)
	return track, err
}

func (self *Client) GetTrackInto(trackId int, mapTo any) error {
	err := self.get("track", strconv.Itoa(trackId), mapTo)
	if apiErr, isApiErr := err.(*qobuzsdk.ApiError); isApiErr {
		if apiErr.Code == 404 {
			err = &qobuzsdk.TrackNotFoundError{TrackId: trackId, Err: err}
		}
	}
	return err
}

func (self *Client) SearchTracks(
	query string,
	limit int16,
	offset int16,
) (qobuzsdk.TracksSearchResult, error) {
	var searchResult qobuzsdk.TracksSearchResult
	err := self.SearchTracksInto(query, limit, offset, &searchResult)
	return searchResult, err
}

func (self *Client) SearchTracksInto(
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.search("track", query, limit, offset, mapTo)
}

func (self *Client) StreamTrack(trackId int, format qobuzsdk.Format) (qobuzsdk.TrackFileUrl, error) {
	var trackFileUrl qobuzsdk.TrackFileUrl
	err := self.StreamTrackInto(trackId, format, &trackFileUrl)
	return trackFileUrl, err
}

func (self *Client) StreamTrackInto(trackId int, format qobuzsdk.Format, mapTo any) error {
	var timestamp int64 = self.getCurrentTimestamp()

	sum := md5.Sum(fmt.Appendf(nil,
		"trackgetFileUrlformat_id%dintentstreamtrack_id%d%d%s",
		format,
		trackId,
		timestamp,
		self.appSecret,
	))
	requestSig := hex.EncodeToString(sum[:])

	endpoint := fmt.Sprintf(
		"track/getFileUrl?track_id=%d&format_id=%d&intent=stream&request_ts=%d&request_sig=%s",
		trackId,
		format,
		timestamp,
		requestSig,
	)
	err := self.apiClient.Get(endpoint, mapTo)
	if apiErr, isApiErr := err.(*qobuzsdk.ApiError); isApiErr {
		if apiErr.Code == 404 {
			err = &qobuzsdk.TrackNotFoundError{TrackId: trackId, Err: err}
		}
	}
	return err
}

func (self *Client) AddTrackToFavorites(trackId int) error {
	favorite := qobuzsdk.Favorite{TrackIds: []int{trackId}}
	return self.AddToFavorites(favorite)
}

func (self *Client) AddAlbumToFavorites(albumId string) error {
	favorite := qobuzsdk.Favorite{AlbumIds: []string{albumId}}
	return self.AddToFavorites(favorite)
}

func (self *Client) AddArtistToFavorites(artistId int) error {
	favorite := qobuzsdk.Favorite{ArtistIds: []int{artistId}}
	return self.AddToFavorites(favorite)
}

func (self *Client) AddToFavorites(favorite qobuzsdk.Favorite) error {
	return self.apiClient.PostForm("favorite/create", favorite)
}

func (self *Client) RemoveTrackFromFavorites(trackId int) error {
	favorite := qobuzsdk.Favorite{TrackIds: []int{trackId}}
	return self.RemoveFromFavorites(favorite)
}

func (self *Client) RemoveAlbumFromFavorites(albumId string) error {
	favorite := qobuzsdk.Favorite{AlbumIds: []string{albumId}}
	return self.RemoveFromFavorites(favorite)
}

func (self *Client) RemoveArtistFromFavorites(artistId int) error {
	favorite := qobuzsdk.Favorite{ArtistIds: []int{artistId}}
	return self.RemoveFromFavorites(favorite)
}

func (self *Client) RemoveFromFavorites(favorite qobuzsdk.Favorite) error {
	return self.apiClient.PostForm("favorite/delete", favorite)
}

func (self *Client) GetAlbum(albumId string) (qobuzsdk.Album, error) {
	var album qobuzsdk.Album
	err := self.GetAlbumInto(albumId, &album)
	return album, err
}

func (self *Client) GetAlbumInto(albumId string, mapTo any) error {
	err := self.get("album", albumId, mapTo)
	if apiErr, isApiErr := err.(*qobuzsdk.ApiError); isApiErr {
		if apiErr.Code == 404 {
			err = &qobuzsdk.AlbumNotFoundError{AlbumId: albumId, Err: err}
		}
	}
	return err
}

func (self *Client) SearchAlbums(
	query string,
	limit int16,
	offset int16,
) (qobuzsdk.AlbumsSearchResult, error) {
	var searchResult qobuzsdk.AlbumsSearchResult
	err := self.SearchAlbumsInto(query, limit, offset, &searchResult)
	return searchResult, err
}

func (self *Client) SearchAlbumsInto(
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.search("album", query, limit, offset, mapTo)
}

func (self *Client) GetPlaylist(playlistId int) (qobuzsdk.Playlist, error) {
	var playlist qobuzsdk.Playlist
	err := self.GetPlaylistInto(playlistId, &playlist)
	return playlist, err
}

func (self *Client) GetPlaylistInto(playlistId int, mapTo any) error {
	err := self.get("playlist", strconv.Itoa(playlistId), mapTo)
	if apiErr, isApiErr := err.(*qobuzsdk.ApiError); isApiErr {
		if apiErr.Code == 404 {
			err = &qobuzsdk.PlaylistNotFoundError{PlaylistId: playlistId, Err: err}
		}
	}
	return err
}

func (self *Client) SearchPlaylists(
	query string,
	limit int16,
	offset int16,
) (qobuzsdk.PlaylistsSearchResult, error) {
	var searchResult qobuzsdk.PlaylistsSearchResult
	err := self.SearchPlaylistsInto(query, limit, offset, &searchResult)
	return searchResult, err
}

func (self *Client) SearchPlaylistsInto(
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.search("playlist", query, limit, offset, mapTo)
}

func (self *Client) GetArtist(artistId int) (qobuzsdk.ArtistPage, error) {
	var artist qobuzsdk.ArtistPage
	err := self.GetArtistInto(artistId, &artist)
	return artist, err
}

func (self *Client) GetArtistInto(artistId int, mapTo any) error {
	endpoint := fmt.Sprintf(
		"artist/page?artist_id=%d",
		artistId,
	)
	err := self.apiClient.Get(endpoint, mapTo)
	if apiErr, isApiErr := err.(*qobuzsdk.ApiError); isApiErr {
		if apiErr.Code == 404 {
			err = &qobuzsdk.ArtistNotFoundError{ArtistId: artistId, Err: err}
		}
	}
	return err
}

func (self *Client) SearchArtists(
	query string,
	limit int16,
	offset int16,
) (qobuzsdk.ArtistsSearchResult, error) {
	var searchResult qobuzsdk.ArtistsSearchResult
	err := self.SearchArtistsInto(query, limit, offset, &searchResult)
	return searchResult, err
}

func (self *Client) SearchArtistsInto(
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.search("artist", query, limit, offset, mapTo)
}

func (self *Client) SearchCatalog(
	query string,
	limit int16,
	offset int16,
) (qobuzsdk.CatalogSearchResult, error) {
	var searchResult qobuzsdk.CatalogSearchResult
	err := self.SearchCatalogInto(query, limit, offset, &searchResult)
	return searchResult, err
}

func (self *Client) SearchCatalogInto(
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.search("catalog", query, limit, offset, mapTo)
}

func (self *Client) get(resourceName string, id string, mapTo any) error {
	return self.apiClient.Get(fmt.Sprintf(
		"%s/get?%s_id=%s",
		resourceName,
		resourceName,
		id,
	), mapTo)
}

func (self *Client) search(
	resourceName string,
	query string,
	limit int16,
	offset int16,
	mapTo any,
) error {
	return self.apiClient.Get(fmt.Sprintf(
		"%s/search?query=%s&limit=%d&offset=%d",
		resourceName,
		query,
		limit,
		offset,
	), mapTo)
}
