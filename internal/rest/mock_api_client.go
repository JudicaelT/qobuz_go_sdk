package rest

import "fmt"

type MockEndpointHandler func(parseTo any) error

type MockApiClient struct {
	headers       map[string]string
	getEndpoints  map[string]MockEndpointHandler
	postEndpoints map[string]MockEndpointHandler
}

func NewMockApiClient() *MockApiClient {
	return &MockApiClient{
		headers:       make(map[string]string),
		getEndpoints:  make(map[string]MockEndpointHandler),
		postEndpoints: make(map[string]MockEndpointHandler),
	}
}

func (self *MockApiClient) MockGetEndpoint(endpoint string, handle MockEndpointHandler) {
	self.getEndpoints[endpoint] = handle
}

func (self *MockApiClient) MockPostEndpoint(endpoint string, handle MockEndpointHandler) {
	self.postEndpoints[endpoint] = handle
}

func (self *MockApiClient) Get(endpoint string, parseTo any) error {
	handleRequest, isSet := self.getEndpoints[endpoint]
	if !isSet {
		return fmt.Errorf("Endpoint 'GET %s' not found", endpoint)
	}
	return handleRequest(parseTo)
}

func (self *MockApiClient) PostForm(
	endpoint string,
	payload any,
	parseTo any,
) error {
	handleRequest, isSet := self.postEndpoints[endpoint]
	if !isSet {
		return fmt.Errorf("Endpoint 'POST %s' not found", endpoint)
	}
	return handleRequest(parseTo)
}

func (self *MockApiClient) SetHeader(key, value string) {
	self.headers[key] = value
}

func (self *MockApiClient) HasHeader(key string) bool {
	_, isSet := self.headers[key]
	return isSet
}
