package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type AddTracksToPlaylistTestSuite struct {
	*ApiTestSuite[any]
	endpoint   string
	playlistId int
}

func TestAddTracksToPlaylistTestSuite(t *testing.T) {
	playlistId := 42
	suite.Run(t, &AddTracksToPlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		endpoint:     "playlist/addTracks",
		playlistId:   playlistId,
	})
}

func (test *AddTracksToPlaylistTestSuite) TestAddTracksToPlaylist() {
	// GIVEN an API client that accepts adding tracks to a playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN adding tracks to the given playlist
	err := client.AddTracksToPlaylist(test.playlistId, 42, 29)

	// THEN there should be no error
	test.NoError(err)
}

func (test *AddTracksToPlaylistTestSuite) TestAddTracksToPlaylist_AddTracksToPlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/addTracks endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/addTracks failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN adding tracks to the given playlist
	err := client.AddTracksToPlaylist(test.playlistId, 42, 29)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *AddTracksToPlaylistTestSuite) TestAddTracksToPlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/addTracks endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN adding tracks to a playlist
	err := client.AddTracksToPlaylist(test.playlistId, 42, 29)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}
