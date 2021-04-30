package client

import "net/http"

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
