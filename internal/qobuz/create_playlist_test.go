package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type CreatePlaylistTestSuite struct {
	*ApiTestSuite[qobuzsdk.CreatedPlaylist]
	endpoint string
}

func TestCreatePlaylistTestSuite(t *testing.T) {
	expectedPlaylist := qobuzsdk.CreatedPlaylist{
		Id:   42,
		Name: "Some Playlist",
	}
	suite.Run(t, &CreatePlaylistTestSuite{
		ApiTestSuite: NewApiTestSuite(expectedPlaylist),
		endpoint:     "playlist/create",
	})
}

func (test *CreatePlaylistTestSuite) TestCreatePlaylist() {
	// GIVEN an API client that accepts creating a playlist
	client := test.NewClient(test.MockPostApiClient(test.endpoint, nil))

	// WHEN creating a playlist
	createdPlaylist, err := client.CreatePlaylist(
		test.ExpectedResult.IsPublic,
		test.ExpectedResult.IsCollaborative,
		test.ExpectedResult.Name,
		test.ExpectedResult.Description,
	)

	// THEN the created playlist should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, createdPlaylist)
}

func (test *CreatePlaylistTestSuite) TestCreatePlaylist_CreatePlaylistEndpointFails() {
	// GIVEN an API client whose POST playlist/create endpoint fails
	postErr := &qobuzsdk.ApiError{Code: 500, Message: "POST playlist/create failed"}
	client := test.NewClient(test.MockPostApiClient(test.endpoint, postErr))

	// WHEN creating a playlist
	createdPlaylist, err := client.CreatePlaylist(
		test.ExpectedResult.IsPublic,
		test.ExpectedResult.IsCollaborative,
		test.ExpectedResult.Name,
		test.ExpectedResult.Description,
	)

	// THEN the endpoint error is returned
	test.ErrorIs(err, postErr)
	// AND no playlist data is returned
	test.Equal(qobuzsdk.CreatedPlaylist{}, createdPlaylist)
}
