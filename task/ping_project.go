package task

import "encoding/json"

// PingProjectConfiguration ...
type PingProjectConfiguration struct {
	AdpProgressTaskTimeout       int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled            bool   `json:"adp_loggingEnabled,omitempty"`
	AdpPingProjectPassword       string `json:"adp_pingProject_Password,omitempty"`
	AdpTaskActive                bool   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout               int    `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent       bool   `json:"adp_executionPersistent,omitempty"`
	AdpPingProjectUser           string `json:"adp_pingProject_User,omitempty"`
	AdpAbortWfOnFailure          bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpPingProjectIdentifierType string `json:"adp_pingProject_IdentifierType,omitempty"`
	AdpPingProjectJSONOutput     string `json:"adp_pingProject_JsonOutput,omitempty"`
	AdpCleanUpHistory            bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpPingProjectIdentifiers    string `json:"adp_pingProject_Identifiers,omitempty"`
}

// EnableAdpLogging ...
func (c *PingProjectConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

// EnableAdpExecutionPersistent ...
func (c *PingProjectConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

// NewPingProjectTaskRequest ...
func NewPingProjectTaskRequest(opts ...func(*PingProjectConfiguration)) *Request {
	cfg := &PingProjectConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Ping Project",
		TaskDescription:   "Pings applications or engines",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Ping Project Task",
	}
}

// WithPingProjectIdentifiers ...
func WithPingProjectIdentifiers(s string) func(*PingProjectConfiguration) {
	return func(c *PingProjectConfiguration) {
		c.AdpPingProjectIdentifiers = s
	}
}

// PingProjectExecutionMetaData ...
type PingProjectExecutionMetaData struct {
	PingProjectResult json.RawMessage `json:"ping_project_result"`
}

// Output ...
func (meta *PingProjectExecutionMetaData) Output() string {

	output := string(meta.PingProjectResult)
	unquoteJSONOutput(&output)

	return output
}

func NewPingProjectTaskResponse() *Response {
	return &Response{
		ExecutionMetaData: &PingProjectExecutionMetaData{},
	}
}
