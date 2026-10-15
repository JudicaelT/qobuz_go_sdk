package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type UnsubscribeFromPlaylistTestSuite struct {
	*ApiTestSuite[any]
	playlistId int
	endpoint   string
}

func TestUnsubscribeFromPlaylistTestSuite(t *testing.T) {
	suite.Run(t, &UnsubscribeFromPlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		playlistId:   42,
		endpoint:     "playlist/unsubscribe",
	})
}

func (test *UnsubscribeFromPlaylistTestSuite) TestUnsubscribeFromPlaylist() {
	// GIVEN an API client that accepts unsubscribing from the given playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN unsubscribing to the playlist
	err := client.UnsubscribeFromPlaylist(test.playlistId)

	// THEN there should be no error
	test.NoError(err)
}

func (test *UnsubscribeFromPlaylistTestSuite) TestUnsubscribeFromPlaylist_UnsubscribeFromPlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/unsubscribe endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/unsubscribe failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN unsubscribing from the playlist
	err := client.UnsubscribeFromPlaylist(test.playlistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *UnsubscribeFromPlaylistTestSuite) TestUnsubscribeFromPlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/unsubscribe endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN unsubscribing from the playlist
	err := client.UnsubscribeFromPlaylist(test.playlistId)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}
