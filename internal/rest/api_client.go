package rest

import (
	"encoding/json"
	"io"
	"net/http"
	urlPackage "net/url"
	"strings"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/gorilla/schema"
)

type ApiClient interface {
	Get(endpoint string, parseTo any) error
	PostForm(endpoint string, payload any, parseTo any) error
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

var encoder *schema.Encoder = schema.NewEncoder()

func (self *apiClientImpl) Get(endpoint string, parseTo any) error {
	res, err := self.call(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := self.assertResIsSuccessful(res); err != nil {
		return err
	}

	if err := self.parseBody(res, parseTo); err != nil {
		return err
	}
	return err
}

func (self *apiClientImpl) PostForm(
	endpoint string,
	payload any,
	parseTo any,
) error {
	res, err := self.call(http.MethodPost, endpoint, payload)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := self.assertResIsSuccessful(res); err != nil {
		return err
	}
	if parseTo != nil {
		if err := self.parseBody(res, parseTo); err != nil {
			return err
		}
	}
	return err
}

func (self *apiClientImpl) call(
	httpMethod string,
	endpoint string,
	payload any,
) (*http.Response, error) {
	url, err := urlPackage.Parse(self.baseUrl + "/" + endpoint)
	if err != nil {
		return nil, err
	}
	url.RawQuery = url.Query().Encode()

	var body io.Reader
	if payload != nil {
		form := urlPackage.Values{}
		if err := encoder.Encode(payload, form); err != nil {
			return nil, err
		}
		body = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(httpMethod, url.String(), body)
	if err != nil {
		return nil, err
	}
	if httpMethod == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for key, value := range self.headers {
		req.Header.Add(key, value)
	}
	return self.httpClient.Do(req)
}

func (self *apiClientImpl) assertResIsSuccessful(res *http.Response) error {
	if res.StatusCode < 200 || res.StatusCode > 299 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		apiErr := &qobuzsdk.ApiError{}
		if err := json.Unmarshal(body, apiErr); err != nil {
			return err
		}
		return apiErr
	}
	return nil
}

func (self *apiClientImpl) parseBody(res *http.Response, parseTo any) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, parseTo)
}

func (self *apiClientImpl) SetHeader(key, value string) {
	self.headers[key] = value
}

func (self *apiClientImpl) HasHeader(key string) bool {
	_, isSet := self.headers[key]
	return isSet
}
