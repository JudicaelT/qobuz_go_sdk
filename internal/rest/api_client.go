package rest

import (
	"encoding/json"
	"io"
	"net/http"
	urlPackage "net/url"

	"github.com/JudicaelT/qobuzsdk"
)

type ApiClient interface {
	Get(endpoint string, parseTo any) error
	SetHeader(key, value string)
	HasHeader(key string) bool
}

type apiClientImpl struct {
	baseUrl    string
	httpClient *http.Client
	headers    map[string]string
}

func NewApiClient(baseUrl string) ApiClient {
	return &apiClientImpl{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
		headers:    make(map[string]string),
	}
}

func (self *apiClientImpl) Get(endpoint string, parseTo any) error {
	url, err := self.buildUrl(endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for key, value := range self.headers {
		req.Header.Add(key, value)
	}

	res, err := self.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return self.parseBody(res, parseTo)
}

func (self *apiClientImpl) buildUrl(endpoint string) (string, error) {
	u, err := urlPackage.Parse(self.baseUrl + "/" + endpoint)
	if err != nil {
		return "", err
	}
	u.RawQuery = u.Query().Encode()
	return u.String(), nil
}

func (self *apiClientImpl) parseBody(res *http.Response, parseTo any) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		apiErr := &qobuzsdk.ApiError{}
		if err := json.Unmarshal(body, apiErr); err != nil {
			return err
		}
		return apiErr
	}
	if err := json.Unmarshal(body, parseTo); err != nil {
		return err
	}
	return nil
}

func (self *apiClientImpl) SetHeader(key, value string) {
	self.headers[key] = value
}

func (self *apiClientImpl) HasHeader(key string) bool {
	_, isSet := self.headers[key]
	return isSet
}
