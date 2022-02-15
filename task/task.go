package task

import (
	"encoding/json"
)

type TaskConfigurator interface {
	EnableAdpLogging()
	EnableAdpExecutionPersistent()
}

type TaskRequest struct {
	TaskType          string           `json:"taskType"`
	TaskDescription   string           `json:"taskDescription"`
	TaskDisplayName   string           `json:"taskDisplayName"`
	TaskConfiguration TaskConfigurator `json:"taskConfiguration"`
}

func (req *TaskRequest) Type(s string) *TaskRequest {
	req.TaskType = s
	return req
}

func (req *TaskRequest) DisplayName(s string) *TaskRequest {
	req.TaskDisplayName = s
	return req
}

func (req *TaskRequest) Description(s string) *TaskRequest {
	req.TaskDescription = s
	return req
}

func (req *TaskRequest) JSON() string {
	b, _ := json.MarshalIndent(req, "", "    ")
	return string(b)
}

type MetaData interface {
	Output() string
}

type TaskResponse struct {
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
