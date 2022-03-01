package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"opentext.com/axcelerate/adp/task"
)

var ADP ADPClient

type ADPClient struct {
	TaskReq  *task.Request
	TaskResp *task.Response
	RC       *RestClient
}

func (c *ADPClient) NewRestRequest() (*http.Request, error) {
	payload, err := json.Marshal(c.TaskReq)
	if err != nil {
		return nil, err
	}

	fmt.Printf("=== payload === %s\n", payload)

	// contruct rest request
	req, err := http.NewRequest(http.MethodPut, c.RC.endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-Username", c.RC.user)
	req.Header.Set("Auth-Password", c.RC.password)

	return req, nil
}

func (c *ADPClient) Run() error {
	var req *http.Request
	var resp *http.Response

	var data []byte
	var err error

	if req, err = c.NewRestRequest(); err != nil {
		return err
	}

	if resp, err = c.RC.httpClient.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	fmt.Printf("=== output === %s\n", data)

	if err = json.Unmarshal(data, c.TaskResp); err != nil {
		return err
	}

	return err
}
