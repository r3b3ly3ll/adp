package task

import "encoding/json"

// StartApplicationConfiguration ...
type StartApplicationConfiguration struct {
	AdpProgressTaskTimeout                   int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                        bool   `json:"adp_loggingEnabled,omitempty"`
	AdpTaskActive                            bool   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                           int    `json:"adp_taskTimeout,omitempty"`
	AdpStartApplicationUseHTTPS              bool   `json:"adp_startApplication_useHttps,omitempty"`
	AdpStartApplicationApplicationURL        string `json:"adp_startApplication_applicationUrl,omitempty"`
	AdpExecutionPersistent                   bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                      bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                        bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpStartApplicationApplicationIdentifier string `json:"adp_startApplication_applicationIdentifier,omitempty"`
}

// EnableAdpLogging ...
func (c *StartApplicationConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

// EnableAdpExecutionPersistent ...
func (c *StartApplicationConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

// NewStartApplicationTaskRequest ...
func NewStartApplicationTaskRequest(opts ...func(*StartApplicationConfiguration)) *Request {
	cfg := &StartApplicationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Start Application",
		TaskDescription:   "Starts an application",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Start Application",
	}
}

// WithStartApplicationApplicationIdentifier ...
func WithStartApplicationApplicationIdentifier(s string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationIdentifier = s
	}
}

// WithStartApplicationApplicationURL ...
func WithStartApplicationApplicationURL(s string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationURL = s
	}
}

// StartApplicationExecutionMetaData ...
type StartApplicationExecutionMetaData json.RawMessage

// Output ...
func (meta *StartApplicationExecutionMetaData) Output() string {
	return string(*meta)
}

func NewStartApplictionTaskResponse() *Response {
	return &Response{
		ExecutionMetaData: &StartApplicationExecutionMetaData{},
	}
}
