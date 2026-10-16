package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	urlPackage "net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/JudicaelT/qobuzsdk"
)

type ApiClient interface {
	Get(endpoint string, parseTo any) error
	PostForm(endpoint string, payload any, parseTo any) error
	SetHeader(key, value string)
	HasHeader(key string) bool
}

type apiClientImpl struct {
	baseUrl     string
	httpClient  *http.Client
	headersLock sync.Mutex
	headers     map[string]string
}

func NewApiClient(baseUrl string) ApiClient {
	return &apiClientImpl{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
		headers:    make(map[string]string),
	}
}

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
		formValues, err := self.encodePayload(payload)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(formValues.Encode())
	}

	req, err := http.NewRequest(httpMethod, url.String(), body)
	if err != nil {
		return nil, err
	}
	if httpMethod == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	self.headersLock.Lock()
	for key, value := range self.headers {
		req.Header.Add(key, value)
	}
	self.headersLock.Unlock()
	return self.httpClient.Do(req)
}

func (self *apiClientImpl) encodePayload(payload any) (url.Values, error) {
	values := url.Values{}
	reflectValue := reflect.ValueOf(payload)
	reflectType := reflectValue.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		if !field.IsExported() {
			return nil, fmt.Errorf("Payload contains unexported field %s", field.Name)
		}

		paramName, _, _ := strings.Cut(field.Tag.Get("form"), ",")
		if paramName == "" {
			return nil, fmt.Errorf("Payload field %s is missing a name", field.Name)
		}

		paramValue, err := self.encodePayloadValue(reflectValue.Field(i))
		if err != nil {
			return nil, fmt.Errorf("Could not encode payload field %s: %w", field.Name, err)
		}
		values.Set(paramName, paramValue)
	}
	return values, nil
}

func (self *apiClientImpl) encodePayloadValue(value reflect.Value) (string, error) {
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil

	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil

	case reflect.Slice:
		parts := make([]string, value.Len())
		for i := 0; i < value.Len(); i++ {
			part, err := self.encodePayloadValue(value.Index(i))
			if err != nil {
				return "", err
			}
			parts[i] = part
		}
		return strings.Join(parts, ","), nil

	default:
		return "", fmt.Errorf("Unsupported type %s", value.Type())
	}
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
	self.headersLock.Lock()
	defer self.headersLock.Unlock()
	self.headers[key] = value
}

func (self *apiClientImpl) HasHeader(key string) bool {
	self.headersLock.Lock()
	defer self.headersLock.Unlock()
	_, isSet := self.headers[key]
	return isSet
}
