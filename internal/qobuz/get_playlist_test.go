package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type GetPlaylistTestSuite struct {
	*ApiTestSuite[qobuzsdk.Playlist]
	playlistId int
	endpoint   string
}

func TestGetPlaylistTestSuite(t *testing.T) {
	playlistId := 42
	expectedPlaylist := qobuzsdk.Playlist{Id: playlistId, Name: "Some Playlist"}
	testSuite := NewApiTestSuite(expectedPlaylist)
	endpoint := fmt.Sprintf(
		"playlist/get?playlist_id=%d",
		playlistId,
	)
	suite.Run(t, &GetPlaylistTestSuite{
		ApiTestSuite: testSuite,
		playlistId:   playlistId,
		endpoint:     endpoint,
	})
}

func (test *GetPlaylistTestSuite) TestGetPlaylist() {
	// GIVEN an API client that returns the playlist for the given id
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the playlist
	result, err := client.GetPlaylist(test.playlistId)

	// THEN the expected playlist should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, result)
}

func (test *GetPlaylistTestSuite) TestGetPlaylist_GetPlaylistEndpointFails() {
	// GIVEN an API client whose GET playlist/get endpoint fails
	getErr := &qobuzsdk.ApiError{Code: 500, Message: "GET playlist/get failed"}
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, getErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the playlist
	result, err := client.GetPlaylist(test.playlistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, getErr)
	// AND no playlist data is returned
	test.Equal(qobuzsdk.Playlist{}, result)
}

func (test *GetPlaylistTestSuite) TestGetPlaylist_PlaylistNotFound() {
	// GIVEN an API client whose GET playlist/get endpoint returns a 404
	errorMessage := "Playlist not found"
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, &qobuzsdk.ApiError{Code: 404, Message: errorMessage}),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the playlist
	result, err := client.GetPlaylist(test.playlistId)

	// THEN a PlaylistNotFoundError is returned
	var notFoundErr *qobuzsdk.PlaylistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.playlistId, notFoundErr.PlaylistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
	// AND no playlist data is returned
	test.Equal(qobuzsdk.Playlist{}, result)
}
