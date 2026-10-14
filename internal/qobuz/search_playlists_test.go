package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type SearchPlaylistsTestSuite struct {
	*ApiTestSuite[qobuzsdk.PlaylistsSearchResult]
	query    string
	limit    int16
	offset   int16
	endpoint string
}

func TestSearchPlaylistsTestSuite(t *testing.T) {
	query := "test query"
	limit := int16(10)
	offset := int16(0)
	expectedSearchResult := qobuzsdk.PlaylistsSearchResult{
		Query:     query,
		Playlists: qobuzsdk.Pagination[qobuzsdk.Playlist]{},
	}
	testSuite := NewApiTestSuite(expectedSearchResult)
	endpoint := fmt.Sprintf(
		"playlist/search?query=%s&limit=%d&offset=%d",
		query,
		limit,
		offset,
	)
	suite.Run(t, &SearchPlaylistsTestSuite{
		ApiTestSuite: testSuite,
		query:        query,
		limit:        limit,
		offset:       offset,
		endpoint:     endpoint,
	})
}

func (test *SearchPlaylistsTestSuite) TestSearchPlaylists() {
	// GIVEN an API client that returns the playlists for the given query
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the playlists
	searchResult, err := client.SearchPlaylists(test.query, test.limit, test.offset)

	// THEN the expected search results should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, searchResult)
}

func (test *SearchPlaylistsTestSuite) TestSearchPlaylists_SearchPlaylistsEndpointFails() {
	// GIVEN an API client whose GET playlist/search endpoint fails
	searchErr := &qobuzsdk.ApiError{Code: 500, Message: "GET playlist/search failed"}
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, searchErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the playlists
	searchResult, err := client.SearchPlaylists(test.query, test.limit, test.offset)

	// THEN the endpoint error is returned
	test.ErrorIs(err, searchErr)
	// AND no search results are returned
	test.Equal(qobuzsdk.PlaylistsSearchResult{}, searchResult)
}
