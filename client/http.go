package client

import (
	"net/http"
	"time"
)


type HTTPClient struct {
	client *http.Client
	URI    string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		client: &http.Client{},
		URI:    uri,
	}
}

func (h HTTPClient) Create(title, message string, duratio time.Duration) ([]byte, error) {
	res := []byte(`response for create reminder`)
	return res, nil
}

func (h HTTPClient) Edit(id, title, message string, duratio time.Duration) ([]byte, error) {
	res := []byte(`response for edit reminder`)
	return res, nil
}

func (h HTTPClient) Fetch(ids []string) ([]byte, error) {
	res := []byte(`response for fetch reminder`)
	return res, nil
}

func (h HTTPClient) Delete(ids []string) error {
	return nil
}

func (h HTTPClient) HealthCheck(host string) bool {
	return true
}

