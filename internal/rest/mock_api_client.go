package rest

import "fmt"

type mockEndpoint struct {
	expectedHeaders map[string]string
	handle          func(parseTo any) error
}

type MockApiClient struct {
	headers   map[string]string
	endpoints map[string]mockEndpoint
}

func NewMockApiClient() *MockApiClient {
	return &MockApiClient{
		headers:   make(map[string]string),
		endpoints: make(map[string]mockEndpoint),
	}
}

func (self *MockApiClient) MockEndpoint(
	endpoint string,
	handle func(parseTo any) error,
	expectedHeaders map[string]string,
) {
	if expectedHeaders == nil {
		expectedHeaders = make(map[string]string)
	}
	self.endpoints[endpoint] = mockEndpoint{
		expectedHeaders: expectedHeaders,
		handle:          handle,
	}
}

func (self *MockApiClient) Get(endpoint string, parseTo any) error {
	mockEndpoint, isSet := self.endpoints[endpoint]
	if !isSet {
		return fmt.Errorf("Endpoint '%s' not found", endpoint)
	}
	for key, expectedValue := range mockEndpoint.expectedHeaders {
		actualValue, headerIsSet := self.headers[key]
		if !headerIsSet {
			return fmt.Errorf(
				"Header '%s' is required to use enpoint '%s'",
				key,
				endpoint,
			)
		}
		if expectedValue != actualValue {
			return fmt.Errorf(
				"Value '%s' of header '%s' should have been '%s'",
				actualValue,
				key,
				expectedValue,
			)
		}
	}
	return mockEndpoint.handle(parseTo)
}

func (self *MockApiClient) SetHeader(key, value string) {
	self.headers[key] = value
}

func (self *MockApiClient) HasHeader(key string) bool {
	_, isSet := self.headers[key]
	return isSet
}
