package qobuz_test

import (
	"fmt"
	"time"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"
	"github.com/JudicaelT/qobuzsdk/internal/rest"

	"github.com/stretchr/testify/suite"
)

type ApiTestSuite[T any] struct {
	suite.Suite
	AppId            int32
	AppSecret        string
	PrivateKey       string
	CurrentTimestamp int64
	ExpectedResult   T
}

func NewApiTestSuite[T any](expectedResult T) *ApiTestSuite[T] {
	return &ApiTestSuite[T]{
		AppId:            123456789,
		AppSecret:        "app_secret",
		PrivateKey:       "private_key",
		CurrentTimestamp: time.Now().Unix(),
		ExpectedResult:   expectedResult,
	}
}

func (test *ApiTestSuite[T]) NewClient(mockApiClient *rest.MockApiClient) qobuzsdk.Client {
	return qobuz.NewClient(
		mockApiClient,
		test.AppId,
		test.AppSecret,
		test.PrivateKey,
		test.GetCurrentTimestamp,
	)
}

func (test *ApiTestSuite[T]) GetCurrentTimestamp() int64 {
	return test.CurrentTimestamp
}

func (test *ApiTestSuite[T]) MockGetApiClient(
	endpoint string,
	errToReturn error,
) *rest.MockApiClient {
	var handle rest.GetEndpointHandler
	if errToReturn == nil {
		handle = func(parseTo any) error {
			parsed, isOfExpectedType := parseTo.(*T)
			if !isOfExpectedType {
				return fmt.Errorf(
					"Parameter 'parseTo' should have been of type *%t for endpoint '%s'",
					parseTo,
					endpoint,
				)
			}
			*parsed = test.ExpectedResult
			return nil
		}
	} else {
		handle = func(parseTo any) error { return errToReturn }
	}

	mockApiClient := rest.NewMockApiClient()
	mockApiClient.MockGetEndpoint(endpoint, handle)
	return mockApiClient
}

func (test *ApiTestSuite[T]) MockPostApiClient(
	endpoint string,
	errToReturn error,
) *rest.MockApiClient {
	mockApiClient := rest.NewMockApiClient()
	mockApiClient.MockPostEndpoint(
		endpoint,
		func() error { return errToReturn },
	)
	return mockApiClient
}
