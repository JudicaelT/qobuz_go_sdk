package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type AddToFavoritesTestSuite struct {
	*ApiTestSuite[any]
	favoritingErr error
	workingClient qobuzsdk.Client
	failingClient qobuzsdk.Client
}

func TestAddToFavoritesTestSuite(t *testing.T) {
	apiTestSuite := NewApiTestSuite[any](nil)
	endpoint := "favorite/create"
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/create failed"}

	suite.Run(t, &AddToFavoritesTestSuite{
		ApiTestSuite:  apiTestSuite,
		favoritingErr: favoritingErr,
		workingClient: apiTestSuite.NewClient(apiTestSuite.MockPostApiClient(endpoint, nil)),
		failingClient: apiTestSuite.NewClient(apiTestSuite.MockPostApiClient(endpoint, favoritingErr)),
	})
}

func (test *AddToFavoritesTestSuite) TestAddTrackToFavorites() {
	err := test.workingClient.AddTrackToFavorites(42)
	test.NoError(err)
}

func (test *AddToFavoritesTestSuite) TestAddTrackToFavorites_CreateFavoriteFails() {
	err := test.failingClient.AddTrackToFavorites(42)
	test.ErrorIs(err, test.favoritingErr)
}

func (test *AddToFavoritesTestSuite) TestAddAlbumToFavorites() {
	err := test.workingClient.AddAlbumToFavorites("42")
	test.NoError(err)
}

func (test *AddToFavoritesTestSuite) TestAddAlbumToFavorites_CreateFavoriteFails() {
	err := test.failingClient.AddAlbumToFavorites("42")
	test.ErrorIs(err, test.favoritingErr)
}

func (test *AddToFavoritesTestSuite) TestAddArtistToFavorites() {
	err := test.workingClient.AddArtistToFavorites(42)
	test.NoError(err)
}

func (test *AddToFavoritesTestSuite) TestAddArtistToFavorites_CreateFavoriteFails() {
	err := test.failingClient.AddArtistToFavorites(42)
	test.ErrorIs(err, test.favoritingErr)
}
