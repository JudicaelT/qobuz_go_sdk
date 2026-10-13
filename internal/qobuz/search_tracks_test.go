package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type SearchTracksTestSuite struct {
	*ApiTestSuite[qobuzsdk.TracksSearchResult]
	query    string
	limit    int16
	offset   int16
	endpoint string
}

func TestSearchTracksTestSuite(t *testing.T) {
	query := "test query"
	limit := int16(10)
	offset := int16(0)
	expectedSearchResult := qobuzsdk.TracksSearchResult{
		Query:  query,
		Tracks: qobuzsdk.Pagination[qobuzsdk.Track]{},
	}
	testSuite := NewApiTestSuite(expectedSearchResult)
	endpoint := fmt.Sprintf(
		"track/search?query=%s&limit=%d&offset=%d",
		query,
		limit,
		offset,
	)
	suite.Run(t, &SearchTracksTestSuite{
		ApiTestSuite: testSuite,
		query:        query,
		limit:        limit,
		offset:       offset,
		endpoint:     endpoint,
	})
}

func (test *SearchTracksTestSuite) TestSearchTracks() {
	// GIVEN an API client that returns the tracks for the given query
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the tracks
	searchResult, err := client.SearchTracks(test.query, test.limit, test.offset)

	// THEN the expected search results should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, searchResult)
}

func (test *SearchTracksTestSuite) TestSearchTracks_SearchTracksEndpointFails() {
	// GIVEN an API client whose GET track/search endpoint fails
	searchErr := &qobuzsdk.ApiError{Code: 500, Message: "GET track/search failed"}
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, searchErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the tracks
	searchResult, err := client.SearchTracks(test.query, test.limit, test.offset)

	// THEN the endpoint error is returned
	test.ErrorIs(err, searchErr)
	// AND no search results are returned
	test.Equal(qobuzsdk.TracksSearchResult{}, searchResult)
}
