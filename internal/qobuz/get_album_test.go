package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type GetAlbumTestSuite struct {
	*ApiTestSuite[qobuzsdk.Album]
	albumId  string
	endpoint string
}

func TestGetAlbumTestSuite(t *testing.T) {
	albumId := "some-album-id"
	expectedAlbum := qobuzsdk.Album{Id: albumId, Title: "Some Album"}
	testSuite := NewApiTestSuite(expectedAlbum)
	endpoint := fmt.Sprintf(
		"album/get?album_id=%s",
		albumId,
	)
	suite.Run(t, &GetAlbumTestSuite{
		ApiTestSuite: testSuite,
		albumId:      albumId,
		endpoint:     endpoint,
	})
}

func (test *GetAlbumTestSuite) TestGetAlbum() {
	// GIVEN an API client that returns the album for the given id
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the album
	result, err := client.GetAlbum(test.albumId)

	// THEN the expected album should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, result)
}

func (test *GetAlbumTestSuite) TestGetAlbum_GetAlbumEndpointFails() {
	// GIVEN an API client whose GET album/get endpoint fails
	getErr := &qobuzsdk.ApiError{Code: 500, Message: "GET album/get failed"}
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, getErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the album
	result, err := client.GetAlbum(test.albumId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, getErr)
	// AND no album data is returned
	test.Equal(qobuzsdk.Album{}, result)
}

func (test *GetAlbumTestSuite) TestGetAlbum_AlbumNotFound() {
	// GIVEN an API client whose GET album/get endpoint returns a 404
	errorMessage := "Album not found"
	client := qobuz.NewClient(
		test.MockApiClient(test.endpoint, &qobuzsdk.ApiError{Code: 404, Message: errorMessage}),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the album
	result, err := client.GetAlbum(test.albumId)

	// THEN a AlbumNotFoundError is returned
	var notFoundErr *qobuzsdk.AlbumNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.albumId, notFoundErr.AlbumId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
	// AND no album data is returned
	test.Equal(qobuzsdk.Album{}, result)
}
