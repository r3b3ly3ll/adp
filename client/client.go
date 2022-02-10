package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"

	"opentext.com/axcelerate/adp/task"
)

// Client represents HTTP client to call ADP Rest API.
type Client struct {
	endpoint      string
	user          string
	password      string
	taskAccessKey string
	httpClient    *http.Client
}

// NewClient is constructor of Client
func NewClient(opts ...func(*Client)) *Client {
	cfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &Client{
		httpClient: &http.Client{Transport: cfg},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithEndPoint sets Endpoint field in Client
func WithEndPoint(s string) func(*Client) {
	return func(c *Client) {
		c.endpoint = s
	}
}

// WithUser sets User field in Client
func WithUser(s string) func(*Client) {
	return func(c *Client) {
		c.user = s
	}
}

// WithPassword sets Password field in Client
func WithPassword(s string) func(*Client) {
	return func(c *Client) {
		c.password = s
	}
}

// WithTaskAccessKey sets TaskAccessKey field in Client
func WithTaskAccessKey(s string) func(*Client) {
	return func(c *Client) {
		c.taskAccessKey = s
	}
}

func (c *Client) Do(taskReq *task.TaskRequest) (task.TaskResponse, error) {
	var taskResp task.TaskResponse
	var err error

	payload, err := json.Marshal(*taskReq)
	if err != nil {
		return taskResp, err
	}

	req, _ := http.NewRequest(http.MethodPut, c.endpoint, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-Username", c.user)
	req.Header.Set("Auth-Password", c.password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return taskResp, err
	}
	defer resp.Body.Close()

	return taskResp, nil
}
