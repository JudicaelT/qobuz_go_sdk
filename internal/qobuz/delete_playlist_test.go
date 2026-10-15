package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type DeletePlaylistTestSuite struct {
	*ApiTestSuite[any]
	endpoint   string
	playlistId int
}

func TestDeletePlaylistTestSuite(t *testing.T) {
	suite.Run(t, &DeletePlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		endpoint:     "playlist/delete",
		playlistId:   42,
	})
}

func (test *DeletePlaylistTestSuite) TestDeletePlaylist() {
	// GIVEN an API client that accepts deleting a playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN deleting a playlist
	err := client.DeletePlaylist(test.playlistId)

	// THEN there should be no error
	test.NoError(err)
}

func (test *DeletePlaylistTestSuite) TestDeletePlaylist_DeletePlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/delete endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/delete failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN deleting a playlist
	err := client.DeletePlaylist(test.playlistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *DeletePlaylistTestSuite) TestDeletePlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/delete endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN deleting a playlist
	err := client.DeletePlaylist(test.playlistId)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}
