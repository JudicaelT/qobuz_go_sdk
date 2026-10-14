package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type AddAlbumToFavoritesTestSuite struct {
	*ApiTestSuite[any]
	albumId  string
	endpoint string
}

func TestAddAlbumToFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &AddAlbumToFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		albumId:      "42",
		endpoint:     "favorite/create",
	})
}

func (test *AddAlbumToFavoritesTestSuite) TestAddAlbumToFavorites() {
	// GIVEN an API client that accepts favoriting the given album
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the album to the favorites
	err := client.AddAlbumToFavorites(test.albumId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *AddAlbumToFavoritesTestSuite) TestAddAlbumToFavorites_CreateFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/create endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/create failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the album to the favorites
	err := client.AddAlbumToFavorites(test.albumId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
