package qobuz_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type StreamTrackTestSuite struct {
	*ApiTestSuite[qobuzsdk.TrackFileUrl]
	trackId  int
	format   qobuzsdk.TrackFormat
	endpoint string
}

func TestStreamTrackTestSuite(t *testing.T) {
	trackId := 42
	expectedTrackFile := qobuzsdk.TrackFileUrl{
		TrackId: trackId,
		URL:     "https://base.url/track/42",
	}
	testSuite := NewApiTestSuite(expectedTrackFile)

	sum := md5.Sum(fmt.Appendf(nil,
		"trackgetFileUrlformat_id%dintentstreamtrack_id%d%d%s",
		qobuzsdk.FormatFLACCD,
		trackId,
		testSuite.CurrentTimestamp,
		testSuite.AppSecret,
	))
	requestSig := hex.EncodeToString(sum[:])

	endpoint := fmt.Sprintf(
		"track/getFileUrl?track_id=%d&format_id=%d&intent=stream&request_ts=%d&request_sig=%s",
		trackId,
		qobuzsdk.FormatFLACCD,
		testSuite.CurrentTimestamp,
		requestSig,
	)
	suite.Run(t, &StreamTrackTestSuite{
		ApiTestSuite: testSuite,
		trackId:      trackId,
		format:       qobuzsdk.FormatFLACCD,
		endpoint:     endpoint,
	})
}

func (test *StreamTrackTestSuite) TestStreamTrack() {
	// GIVEN an API client that returns the file url for the given track id
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN streaming the track
	fileUrl, err := client.StreamTrack(test.trackId, test.format)

	// THEN the expected file url should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, fileUrl)
}

func (test *StreamTrackTestSuite) TestStreamTrack_GetFileUrlEndpointFails() {
	// GIVEN an API client whose GET track/getFileUrl endpoint fails
	streamingErr := &qobuzsdk.ApiError{Code: 500, Message: "GET track/getFileUrl failed"}
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, streamingErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN streaming the track
	fileUrl, err := client.StreamTrack(test.trackId, test.format)

	// THEN the endpoint error is returned
	test.ErrorIs(err, streamingErr)
	// AND no TrackFileUrl data is returned
	test.Equal(qobuzsdk.TrackFileUrl{}, fileUrl)
}

func (test *StreamTrackTestSuite) TestStreamTrack_TrackNotFound() {
	// GIVEN an API client whose GET track/getFileUrl endpoint returns a 404
	errorMessage := "Track not found"
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, &qobuzsdk.ApiError{Code: 404, Message: errorMessage}),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN streaming the track
	fileUrl, err := client.StreamTrack(test.trackId, test.format)

	// THEN a TrackNotFoundError is returned
	var trackNotFoundErr *qobuzsdk.TrackNotFoundError
	if test.ErrorAs(err, &trackNotFoundErr) {
		test.Equal(test.trackId, trackNotFoundErr.TrackId)
		test.Equal(errorMessage, trackNotFoundErr.Error())
	}
	// AND no TrackFileUrl data is returned
	test.Equal(qobuzsdk.TrackFileUrl{}, fileUrl)
}
