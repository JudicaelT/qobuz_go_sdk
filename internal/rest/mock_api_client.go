package rest

import "fmt"

type GetEndpointHandler func(parseTo any) error
type PostEndpointHandler func() error

type MockApiClient struct {
	headers       map[string]string
	getEndpoints  map[string]GetEndpointHandler
	postEndpoints map[string]PostEndpointHandler
}

func NewMockApiClient() *MockApiClient {
	return &MockApiClient{
		headers:       make(map[string]string),
		getEndpoints:  make(map[string]GetEndpointHandler),
		postEndpoints: make(map[string]PostEndpointHandler),
	}
}

func (self *MockApiClient) MockGetEndpoint(endpoint string, handle GetEndpointHandler) {
	self.getEndpoints[endpoint] = handle
}

func (self *MockApiClient) MockPostEndpoint(endpoint string, handle PostEndpointHandler) {
	self.postEndpoints[endpoint] = handle
}

func (self *MockApiClient) Get(endpoint string, parseTo any) error {
	handleRequest, isSet := self.getEndpoints[endpoint]
	if !isSet {
		return fmt.Errorf("Endpoint 'GET %s' not found", endpoint)
	}
	return handleRequest(parseTo)
}

func (self *MockApiClient) PostForm(endpoint string, payload any) error {
	handleRequest, isSet := self.postEndpoints[endpoint]
	if !isSet {
		return fmt.Errorf("Endpoint 'POST %s' not found", endpoint)
	}
	return handleRequest()
}

func (self *MockApiClient) SetHeader(key, value string) {
	self.headers[key] = value
}

func (self *MockApiClient) HasHeader(key string) bool {
	_, isSet := self.headers[key]
	return isSet
}
