package qobuz_test

import (
	"fmt"
	"testing"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"

	"github.com/stretchr/testify/suite"
)

type GetArtistTestSuite struct {
	*ApiTestSuite[qobuzsdk.ArtistPage]
	artistId int
	endpoint string
}

func TestGetArtistTestSuite(t *testing.T) {
	artistId := 42
	expectedArtist := qobuzsdk.ArtistPage{
		Id:   artistId,
		Name: qobuzsdk.ArtistPageDisplayName{Display: "Some Artist"},
	}
	testSuite := NewApiTestSuite(expectedArtist)
	endpoint := fmt.Sprintf(
		"artist/page?artist_id=%d",
		artistId,
	)
	suite.Run(t, &GetArtistTestSuite{
		ApiTestSuite: testSuite,
		artistId:     artistId,
		endpoint:     endpoint,
	})
}

func (test *GetArtistTestSuite) TestGetArtist() {
	// GIVEN an API client that returns the artist for the given id
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, nil),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the artist
	result, err := client.GetArtist(test.artistId)

	// THEN the expected artist should be returned
	test.NoError(err)
	test.Equal(test.ExpectedResult, result)
}

func (test *GetArtistTestSuite) TestGetArtist_GetArtistEndpointFails() {
	// GIVEN an API client whose GET artist/get endpoint fails
	getErr := &qobuzsdk.ApiError{Code: 500, Message: "GET artist/get failed"}
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, getErr),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the artist
	result, err := client.GetArtist(test.artistId)

	// THEN the endpoint error is returned
	test.ErrorIs(err, getErr)
	// AND no artist data is returned
	test.Equal(qobuzsdk.ArtistPage{}, result)
}

func (test *GetArtistTestSuite) TestGetArtist_ArtistNotFound() {
	// GIVEN an API client whose GET artist/get endpoint returns a 404
	errorMessage := "Artist not found"
	client := qobuz.NewClient(
		test.MockGetApiClient(test.endpoint, &qobuzsdk.ApiError{Code: 404, Message: errorMessage}),
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)

	// WHEN getting the artist
	result, err := client.GetArtist(test.artistId)

	// THEN a ArtistNotFoundError is returned
	var notFoundErr *qobuzsdk.ArtistNotFoundError
	if test.ErrorAs(err, &notFoundErr) {
		test.Equal(test.artistId, notFoundErr.ArtistId)
		test.Equal(errorMessage, notFoundErr.Error())
	}
	// AND no artist data is returned
	test.Equal(qobuzsdk.ArtistPage{}, result)
}
