package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type RemoveTracksFromPlaylistTestSuite struct {
	*ApiTestSuite[any]
	endpoint   string
	playlistId int
}

func TestRemoveTracksFromPlaylistTestSuite(t *testing.T) {
	playlistId := 42
	suite.Run(t, &RemoveTracksFromPlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		endpoint:     "playlist/deleteTracks",
		playlistId:   playlistId,
	})
}

func (test *RemoveTracksFromPlaylistTestSuite) TestRemoveTracksFromPlaylist() {
	// GIVEN an API client that accepts removing tracks from a playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN removing tracks from the given playlist
	err := client.RemoveTracksFromPlaylist(test.playlistId, 42, 29)

	// THEN there should be no error
	test.NoError(err)
}

func (test *RemoveTracksFromPlaylistTestSuite) TestRemoveTracksFromPlaylist_RemoveTracksFromPlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/deleteTracks endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/deleteTracks failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN removing tracks from the given playlist
	err := client.RemoveTracksFromPlaylist(test.playlistId, 42, 29)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *RemoveTracksFromPlaylistTestSuite) TestRemoveTracksFromPlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/deleteTracks endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN removing tracks from a playlist
	err := client.RemoveTracksFromPlaylist(test.playlistId, 42, 29)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}
