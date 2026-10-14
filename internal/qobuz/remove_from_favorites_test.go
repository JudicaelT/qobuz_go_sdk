package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"

	"github.com/stretchr/testify/suite"
)

type RemoveFromFavoritesTestSuite struct {
	*ApiTestSuite[any]
	unfavoritingErr error
	workingClient   qobuzsdk.Client
	failingClient   qobuzsdk.Client
}

func TestRemoveFromFavoritesTestSuite(t *testing.T) {
	apiTestSuite := NewApiTestSuite[any](nil)
	endpoint := "favorite/delete"
	unfavoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/delete failed"}

	suite.Run(t, &RemoveFromFavoritesTestSuite{
		ApiTestSuite:    apiTestSuite,
		unfavoritingErr: unfavoritingErr,
		workingClient:   apiTestSuite.NewClient(apiTestSuite.MockPostApiClient(endpoint, nil)),
		failingClient:   apiTestSuite.NewClient(apiTestSuite.MockPostApiClient(endpoint, unfavoritingErr)),
	})
}

func (test *RemoveFromFavoritesTestSuite) TestAddTrackToFavorites() {
	err := test.workingClient.RemoveTrackFromFavorites(42)
	test.NoError(err)
}

func (test *RemoveFromFavoritesTestSuite) TestAddTrackToFavorites_DeleteFavoriteFails() {
	err := test.failingClient.RemoveTrackFromFavorites(42)
	test.ErrorIs(err, test.unfavoritingErr)
}

func (test *RemoveFromFavoritesTestSuite) TestAddAlbumToFavorites() {
	err := test.workingClient.RemoveAlbumFromFavorites("42")
	test.NoError(err)
}

func (test *RemoveFromFavoritesTestSuite) TestAddAlbumToFavorites_DeleteFavoriteFails() {
	err := test.failingClient.RemoveAlbumFromFavorites("42")
	test.ErrorIs(err, test.unfavoritingErr)
}

func (test *RemoveFromFavoritesTestSuite) TestAddArtistToFavorites() {
	err := test.workingClient.RemoveArtistFromFavorites(42)
	test.NoError(err)
}

func (test *RemoveFromFavoritesTestSuite) TestAddArtistToFavorites_DeleteFavoriteFails() {
	err := test.failingClient.RemoveArtistFromFavorites(42)
	test.ErrorIs(err, test.unfavoritingErr)
}
