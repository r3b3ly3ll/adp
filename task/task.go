package task

import (
	"encoding/json"
	"strconv"
)

type TaskConfigurator interface {
	EnableAdpLogging()
	EnableAdpExecutionPersistent()
}

type Request struct {
	TaskType          string           `json:"taskType"`
	TaskDescription   string           `json:"taskDescription"`
	TaskDisplayName   string           `json:"taskDisplayName"`
	TaskConfiguration TaskConfigurator `json:"taskConfiguration"`
}

func (req *Request) Type(s string) *Request {
	req.TaskType = s
	return req
}

func (req *Request) DisplayName(s string) *Request {
	req.TaskDisplayName = s
	return req
}

func (req *Request) Description(s string) *Request {
	req.TaskDescription = s
	return req
}

func (req *Request) JSON() string {
	b, _ := json.MarshalIndent(req, "", "    ")
	return string(b)
}

type MetaData interface {
	Output() string
}

type Response struct {
	ExecutionID         string   `json:"executionId"`
	TaskType            string   `json:"taskType"`
	LoggingEnabled      string   `json:"loggingEnabled"`
	ProgressMax         int      `json:"progressMax"`
	ExecutionStatus     string   `json:"executionStatus"`
	ExecutionRootDir    string   `json:"executionRootDir"`
	ContextID           string   `json:"contextId"`
	ExecutionPersistent string   `json:"executionPersistent"`
	ProgressCurrent     int      `json:"progressCurrent"`
	ProgressPercentage  float64  `json:"progressPercentage"`
	TaskDisplayName     string   `json:"taskDisplayName"`
	ExecutionMetaData   MetaData `json:"executionMetaData"`
}

func (resp *Response) IsSuccess() bool {
	return resp.ExecutionStatus == "success"
}

func unquoteJSONOutput(s *string) {
	unescaped, err := strconv.Unquote(*s)
	if err == nil {
		*s = unescaped
	}
}
