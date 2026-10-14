package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type AddTrackToFavoritesTestSuite struct {
	*ApiTestSuite[any]
	trackId  int
	endpoint string
}

func TestAddTrackToFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &AddTrackToFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		trackId:      42,
		endpoint:     "favorite/create",
	})
}

func (test *AddTrackToFavoritesTestSuite) TestAddTrackToFavorites() {
	// GIVEN an API client that accepts favoriting the given track
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the track to the favorites
	err := client.AddTrackToFavorites(test.trackId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *AddTrackToFavoritesTestSuite) TestAddTrackToFavorites_CreateFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/create endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/create failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN adding the track to the favorites
	err := client.AddTrackToFavorites(test.trackId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
