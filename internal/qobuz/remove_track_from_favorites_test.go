package qobuz_test

import (
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type RemoveTrackFromFavoritesTestSuite struct {
	*ApiTestSuite[any]
	trackId  int
	endpoint string
}

func TestRemoveTrackFromFavoritesTestSuite(t *testing.T) {
	suite.Run(t, &RemoveTrackFromFavoritesTestSuite{
		ApiTestSuite: NewApiTestSuite[any](nil),
		trackId:      42,
		endpoint:     "favorite/delete",
	})
}

func (test *RemoveTrackFromFavoritesTestSuite) TestRemoveTrackFromFavorites() {
	// GIVEN an API client that accepts unfavoriting the given track
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the track from the favorites
	err := client.RemoveTrackFromFavorites(test.trackId)

	// THEN no error should be returned
	test.NoError(err)
}

func (test *RemoveTrackFromFavoritesTestSuite) TestRemoveTrackFromFavorites_DeleteFavoriteEndpointFails() {
	// GIVEN an API client whose POST favorite/delete endpoint fails
	favoritingErr := &qobuzsdk.ApiError{Code: 500, Message: "POST favorite/delete failed"}
	client := qobuz.NewClient(
		test.MockPostApiClient(test.endpoint, favoritingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN removing the track from the favorites
	err := client.RemoveTrackFromFavorites(test.trackId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, favoritingErr)
}
