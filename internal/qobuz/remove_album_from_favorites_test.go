package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type RemoveAlbumFromFavoritesTestSuite struct {
	*ApiTestSuite[any]
	albumId  string
	endpoint string
}

func TestRemoveAlbumFromFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &RemoveAlbumFromFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		albumId:      "42",
		endpoint:     "favorite/delete",
	})
}

func (test *RemoveAlbumFromFavoritesTestSuite) TestRemoveAlbumFromFavorites() {
	// GIVEN an API client that accepts unfavoriting the given album
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the album from the favorites
	err := client.RemoveAlbumFromFavorites(test.albumId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *RemoveAlbumFromFavoritesTestSuite) TestRemoveAlbumFromFavorites_DeleteFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/delete endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/delete failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the album from the favorites
	err := client.RemoveAlbumFromFavorites(test.albumId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
