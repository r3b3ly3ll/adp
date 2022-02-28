package task

import "encoding/json"

// StopProcessesConfiguration ...
type StopProcessesConfiguration struct {
	AdpProgressTaskTimeout           int                    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                bool                   `json:"adp_loggingEnabled,omitempty"`
	AdpStopProcessProcessIdentifiers []ProcessIdentifierArg `json:"adp_stopProcess_processIdentifiers,omitempty"`
	AdpTaskActive                    bool                   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                   int                    `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent           bool                   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure              bool                   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                bool                   `json:"adp_cleanUpHistory,omitempty"`
}

// EnableAdpLogging ...
func (c *StopProcessesConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

// EnableAdpExecutionPersistent ...
func (c *StopProcessesConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

// NewStopProcessesTaskRequest ...
func NewStopProcessesTaskRequest(opts ...func(*StopProcessesConfiguration)) *Request {
	cfg := &StopProcessesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Stop Processes",
		TaskDescription:   "Stops Processes",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Stop Processes",
	}
}

// WithStopProcessProcessProcessIdentifiers ...
// @TaskModelParameter(name="adp_stopProcess_processIdentifiers", mandatory=true)
// @TableDescriptor(columnNames="Process identifier|Stop recursive", columnTypes="String|Boolean", separator="|")
func WithStopProcessProcessProcessIdentifiers(s string) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpStopProcessProcessIdentifiers = parseProcessIdentifierArgs(s)
	}
}

// StopProcessesExecutionMetaData ...
type StopProcessesExecutionMetaData json.RawMessage

// Output ...
func (meta *StopProcessesExecutionMetaData) Output() string {
	return string(*meta)
}

func NewStopProcessesTaskResponse() *Response {
	return &Response{
		ExecutionMetaData: &StopProcessesExecutionMetaData{},
	}
}
