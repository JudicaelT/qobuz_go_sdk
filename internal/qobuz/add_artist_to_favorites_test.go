package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type AddArtistToFavoritesTestSuite struct {
	*ApiTestSuite[any]
	artistId int
	endpoint string
}

func TestAddArtistToFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &AddArtistToFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		artistId:     42,
		endpoint:     "favorite/create",
	})
}

func (test *AddArtistToFavoritesTestSuite) TestAddArtistToFavorites() {
	// GIVEN an API client that accepts favoriting the given artist
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the artist to the favorites
	err := client.AddArtistToFavorites(test.artistId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *AddArtistToFavoritesTestSuite) TestAddArtistToFavorites_CreateFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/create endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/create failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the artist to the favorites
	err := client.AddArtistToFavorites(test.artistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
