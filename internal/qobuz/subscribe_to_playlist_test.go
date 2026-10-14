package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type SubscribeToPlaylistTestSuite struct {
	*ApiTestSuite[qobuzsdk.Playlist]
	playlistId int
	endpoint   string
}

func TestSubscribeToPlaylistTestSuite(t *testing.T) {
	playlistId := 42
	expectedPlaylist := qobuzsdk.Playlist{Id: playlistId, Name: "Some Playlist"}
	suite.Run(t, &SubscribeToPlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite(expectedPlaylist),
		playlistId:   playlistId,
		endpoint:     "playlist/subscribe",
	})
}

func (test *SubscribeToPlaylistTestSuite) TestSubscribeToPlaylist() {
	// GIVEN an API client that accepts subscribing to the given playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN subscribing to the playlist
	err := client.SubscribeToPlaylist(test.playlistId)

	// THEN there should be no error
	test.NoError(err)
}

func (test *SubscribeToPlaylistTestSuite) TestSubscribeToPlaylist_SubscribeToPlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/subscribe endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/subscribe failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN subscribing to the playlist
	err := client.SubscribeToPlaylist(test.playlistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *SubscribeToPlaylistTestSuite) TestSubscribeToPlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/subscribe endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN subscribing to the playlist
	err := client.SubscribeToPlaylist(test.playlistId)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}
