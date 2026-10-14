package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type RemoveArtistFromFavoritesTestSuite struct {
	*ApiTestSuite[any]
	artistId int
	endpoint string
}

func TestRemoveArtistFromFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &RemoveArtistFromFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		artistId:     42,
		endpoint:     "favorite/delete",
	})
}

func (test *RemoveArtistFromFavoritesTestSuite) TestRemoveArtistFromFavorites() {
	// GIVEN an API client that accepts unfavoriting the given artist
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the artist from the favorites
	err := client.RemoveArtistFromFavorites(test.artistId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *RemoveArtistFromFavoritesTestSuite) TestRemoveArtistFromFavorites_DeleteFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/delete endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/delete failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the artist from the favorites
	err := client.RemoveArtistFromFavorites(test.artistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
