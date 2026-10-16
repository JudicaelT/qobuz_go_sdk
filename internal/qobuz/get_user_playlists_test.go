package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type GetUserPlaylistsTestSuite struct {
	*ApiTestSuite[qobuzsdk.UserPlaylists]
	endpoint string
	limit    int16
	offset   int16
}

func TestGetUserPlaylistsTestSuite(t *testing.T) {
	limit := int16(10)
	offset := int16(0)
	expectedUserPlaylist := qobuzsdk.UserPlaylists{
		User: qobuzsdk.User{Id: 42},
		Playlists: qobuzsdk.Pagination[qobuzsdk.Playlist]{
			Limit:  limit,
			Offset: offset,
			Total:  1,
			Items: []qobuzsdk.Playlist{
				qobuzsdk.Playlist{Id: 29},
			},
		},
	}
	endpoint := fmt.Sprintf(
		"playlist/getUserPlaylists?limit=%d&offset=%d",
		limit,
		offset,
	)
	suite.Run(t, &GetUserPlaylistsTestSuite{
		ApiTestSuite: NewApiTestSuite(expectedUserPlaylist),
		endpoint:     endpoint,
		limit:        limit,
		offset:       offset,
	})
}

func (test *GetUserPlaylistsTestSuite) TestGetUserPlaylists() {
	// GIVEN an API client that returns the user's playlists
	client := test.NewClient(test.MockGetApiClient(test.endpoint, nil))

	// WHEN getting the user's playlists
	userPlaylists, err := client.GetUserPlaylists(test.limit, test.offset)

	// THEN the expected result should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, userPlaylists)
}

func (test *GetUserPlaylistsTestSuite) TestGetUserPlaylists_GetUserPlaylistsEndpointFails() {
	// GIVEN an API client whose GET playlist/userPlaylists endpoint fails
	getErr := &qobuzsdk.ApiError{Code: 500, Message: "GET playlist/userPlaylists failed"}
	client := test.NewClient(test.MockGetApiClient(test.endpoint, getErr))

	// WHEN getting the user's playlists
	userPlaylists, err := client.GetUserPlaylists(test.limit, test.offset)

	// THEN the endpoint error is returned
	test.ErrorIs(err, getErr)
	// AND no user playlists data is returned
	test.Equal(qobuzsdk.UserPlaylists{}, userPlaylists)
}
