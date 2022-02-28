package client

import (
	"crypto/tls"
	"net/http"
)

// Client represents HTTP client to call ADP Rest API.
type RestClient struct {
	endpoint      string
	user          string
	password      string
	taskAccessKey string
	httpClient    *http.Client
}

// NewRestClient ...
func NewRestClient(opts ...func(*RestClient)) *RestClient {
	cfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &RestClient{
		httpClient: &http.Client{Transport: cfg},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithEndPoint sets Endpoint field in Client
func WithEndPoint(s string) func(*RestClient) {
	return func(c *RestClient) {
		c.endpoint = s
	}
}

// WithUser sets User field in Client
func WithUser(s string) func(*RestClient) {
	return func(c *RestClient) {
		c.user = s
	}
}

// WithPassword sets Password field in Client
func WithPassword(s string) func(*RestClient) {
	return func(c *RestClient) {
		c.password = s
	}
}

// WithTaskAccessKey sets TaskAccessKey field in Client
func WithTaskAccessKey(s string) func(*RestClient) {
	return func(c *RestClient) {
		c.taskAccessKey = s
	}
}
