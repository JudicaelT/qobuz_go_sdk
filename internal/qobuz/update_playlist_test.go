package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type UpdatePlaylistTestSuite struct {
	*ApiTestSuite[any]
	endpoint   string
	playlistId int
}

func TestUpdatePlaylistTestSuite(t *testing.T) {
	playlistId := 42
	suite.Run(t, &UpdatePlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		endpoint:     "playlist/update",
		playlistId:   playlistId,
	})
}

func (test *UpdatePlaylistTestSuite) TestUpdatePlaylist() {
	// GIVEN an API client that accepts updating a playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN updating a playlist
	err := test.updatePlaylist(client)

	// THEN there should be no error
	test.NoError(err)
}

func (test *UpdatePlaylistTestSuite) TestUpdatePlaylist_UpdatePlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/update endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/update failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN updating a playlist
	err := test.updatePlaylist(client)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
}

func (test *UpdatePlaylistTestSuite) TestUpdatePlaylist_PlaylistNotFound() {
	// GIVEN an API client whose POST playlist/update endpoint returns a 404
	errorMessage := "Playlist not found"
	client := test.NewClient(test.MockPostApiClient(
		test.endpoint,
		&qobuzsdk.ApiError{Code: 404, Message: errorMessage},
	))

	// WHEN updating a playlist
	err := test.updatePlaylist(client)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
}

func (test *UpdatePlaylistTestSuite) updatePlaylist(client qobuzsdk.Client) error {
	isPublic := true
	isCollaborative := true
	name := "test playlist name"
	description := "test playlist description"
	return client.UpdatePlaylist(test.playlistId, isPublic, isCollaborative, name, description)
}
