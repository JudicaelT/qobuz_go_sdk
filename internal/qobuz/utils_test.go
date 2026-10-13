package qobuz_test

import (
	"fmt"
	"time"

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

func (test *ApiTestSuite[T]) GetCurrentTimestamp() int64 {
	return test.CurrentTimestamp
}

func (test *ApiTestSuite[T]) MockApiClient(
	endpoint string,
	errToReturn error,
) *rest.MockApiClient {
	var handle func(parseTo any) error
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
	mockApiClient.MockEndpoint(
		endpoint,
		handle,
		nil,
	)
	return mockApiClient
}
