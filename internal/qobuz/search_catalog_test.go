package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type SearchCatalogTestSuite struct {
	*ApiTestSuite[qobuzsdk.CatalogSearchResult]
	query    string
	limit    int16
	offset   int16
	endpoint string
}

func TestSearchCatalogTestSuite(t *testing.T) {
	query := "test query"
	limit := int16(10)
	offset := int16(0)
	expectedSearchResult := qobuzsdk.CatalogSearchResult{
		Query:     query,
		Tracks:    qobuzsdk.Pagination[qobuzsdk.Track]{},
		Albums:    qobuzsdk.Pagination[qobuzsdk.Album]{},
		Playlists: qobuzsdk.Pagination[qobuzsdk.Playlist]{},
		Artists:   qobuzsdk.Pagination[qobuzsdk.ArtistSearchItem]{},
	}
	testSuite := NewApiTestSuite(expectedSearchResult)
	endpoint := fmt.Sprintf(
		"catalog/search?query=%s&limit=%d&offset=%d",
		query,
		limit,
		offset,
	)
	suite.Run(t, &SearchCatalogTestSuite{
		ApiTestSuite: testSuite,
		query:        query,
		limit:        limit,
		offset:       offset,
		endpoint:     endpoint,
	})
}

func (test *SearchCatalogTestSuite) TestSearchCatalog() {
	// GIVEN an API client that returns the catalog for the given query
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the catalog
	searchResult, err := client.SearchCatalog(test.query, test.limit, test.offset)

	// THEN the expected search results should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, searchResult)
}

func (test *SearchCatalogTestSuite) TestSearchCatalog_SearchCatalogEndpointFails() {
	// GIVEN an API client whose GET catalog/search endpoint fails
	searchErr := &qobuzsdk.ApiError{Code: 500, Message: "GET catalog/search failed"}
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, searchErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN searching the catalog
	searchResult, err := client.SearchCatalog(test.query, test.limit, test.offset)

	// THEN the endpoint error is returned
	test.ErrorIs(err, searchErr)
	// AND no search results are returned
	test.Equal(qobuzsdk.CatalogSearchResult{}, searchResult)
}
