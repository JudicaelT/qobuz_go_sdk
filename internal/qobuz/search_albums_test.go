package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type SearchAlbumsTestSuite struct {
	*ApiTestSuite[qobuzsdk.AlbumsSearchResult]
	query    string
	limit    int16
	offset   int16
	endpoint string
}

func TestSearchAlbumsTestSuite(t *testing.T) {
	query := "test query"
	limit := int16(10)
	offset := int16(0)
	expectedSearchResult := qobuzsdk.AlbumsSearchResult{
		Query:  query,
		Albums: qobuzsdk.Pagination[qobuzsdk.Album]{},
	}
	testSuite := NewApiTestSuite(expectedSearchResult)
	endpoint := fmt.Sprintf(
		"album/search?query=%s&limit=%d&offset=%d",
		query,
		limit,
		offset,
	)
	suite.Run(t, &SearchAlbumsTestSuite{
		ApiTestSuite: testSuite,
		query:        query,
		limit:        limit,
		offset:       offset,
		endpoint:     endpoint,
	})
}

func (test *SearchAlbumsTestSuite) TestSearchAlbums() {
	// GIVEN an API client that returns the albums for the given query
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the albums
	searchResult, err := client.SearchAlbums(test.query, test.limit, test.offset)

	// THEN the expected search results should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, searchResult)
}

func (test *SearchAlbumsTestSuite) TestSearchAlbums_SearchAlbumsEndpointFails() {
	// GIVEN an API client whose GET album/search endpoint fails
	searchErr := &qobuzsdk.ApiError{Code: 500, Message: "GET album/search failed"}
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, searchErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the albums
	searchResult, err := client.SearchAlbums(test.query, test.limit, test.offset)

	// THEN the endpoint error is returned
	test.ErrorIs(err, searchErr)
	// AND no search results are returned
	test.Equal(qobuzsdk.AlbumsSearchResult{}, searchResult)
}
