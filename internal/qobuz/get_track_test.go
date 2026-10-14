package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type GetTrackTestSuite struct {
	*ApiTestSuite[qobuzsdk.Track]
	trackId  int
	endpoint string
}

func TestGetTrackTestSuite(t *testing.T) {
	trackId := 42
	expectedTrack := qobuzsdk.Track{Id: trackId, Title: "Some Track"}
	testSuite := NewApiTestSuite(expectedTrack)
	endpoint := fmt.Sprintf(
		"track/get?track_id=%d",
		trackId,
	)
	suite.Run(t, &GetTrackTestSuite{
		ApiTestSuite: testSuite,
		trackId:      trackId,
		endpoint:     endpoint,
	})
}

func (test *GetTrackTestSuite) TestGetTrack() {
	// GIVEN an API client that returns the track for the given id
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the track
	track, err := client.GetTrack(test.trackId)

	// THEN the expected track should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, track)
}

func (test *GetTrackTestSuite) TestGetTrack_GetTrackEndpointFails() {
	// GIVEN an API client whose GET track/get endpoint fails
	getTrackErr := &qobuzsdk.ApiError{Code: 500, Message: "GET track/get failed"}
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, getTrackErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the track
	track, err := client.GetTrack(test.trackId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, getTrackErr)
	// AND no track data is returned
	test.Equal(qobuzsdk.Track{}, track)
}

func (test *GetTrackTestSuite) TestGetTrack_TrackNotFound() {
	// GIVEN an API client whose GET track/get endpoint returns a 404
	errorMessage := "Track not found"
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, &qobuzsdk.ApiError{Code: 404, Message: errorMessage}),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the track
	track, err := client.GetTrack(test.trackId)

	// THEN a TrackNotFoundError is returned
	var notFoundErr *qobuzsdk.TrackNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.trackId, notFoundErr.TrackId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
	// AND no track data is returned
	test.Equal(qobuzsdk.Track{}, track)
}
