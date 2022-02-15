package task

import (
	"encoding/json"
)

// ListEntitiesConfiguration ...
type ListEntitiesConfiguration struct {
	AdpProgressTaskTimeout                                 int      `json:"adp_progressTaskTimeout,omitempty"`
	AdpListEntitiesFile                                    string   `json:"adp_listEntities_file,omitempty"`
	AdpListEntitiesNumberOfEntities                        string   `json:"adp_listEntities_numberOfEntities,omitempty"`
	AdpListEntitiesAxcRequestTimeoutSeconds                int      `json:"adp_listEntities_axcRequestTimeoutSeconds,omitempty"`
	AdpTaskActive                                          bool     `json:"adp_taskActive,omitempty"`
	AdpListEntitiesUserHasAccess                           string   `json:"adp_listEntities_userHasAccess,omitempty"`
	AdpListEntitiesWhiteList                               string   `json:"adp_listEntities_whiteList,omitempty"`
	AdpExecutionPersistent                                 bool     `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                    bool     `json:"adp_abortWfOnFailure,omitempty"`
	AdpListEntitiesRelatedEntity                           string   `json:"adp_listEntities_relatedEntity,omitempty"`
	AdpListEntitiesWorkspace                               string   `json:"adp_listEntities_workspace,omitempty"`
	AdpLoggingEnabled                                      bool     `json:"adp_loggingEnabled,omitempty"`
	AdpListEntitiesStatus                                  string   `json:"adp_listEntities_status,omitempty"`
	AdpListEntitiesAxcServiceCoreAddress                   string   `json:"adp_listEntities_axcServiceCoreAddress,omitempty"`
	AdpListEntitiesRelatedEntityType                       string   `json:"adp_listEntities_relatedEntityType,omitempty"`
	AdpListEntitiesType                                    string   `json:"adp_listEntities_type,omitempty"`
	AdpListEntitiesHTTPSKeystoreFile                       string   `json:"adp_listEntities_httpsKeystoreFile,omitempty"`
	AdpListEntitiesHTTPSPassword                           string   `json:"adp_listEntities_httpsPassword,omitempty"`
	AdpListEntitiesAxcConnectTimeoutSeconds                int      `json:"adp_listEntities_axcConnectTimeoutSeconds,omitempty"`
	AdpListEntitiesAxcServicePassword                      string   `json:"adp_listEntities_axcServicePassword,omitempty"`
	AdpListEntitiesStartingEntity                          string   `json:"adp_listEntities_startingEntity,omitempty"`
	AdpListEntitiesOutputJSON                              string   `json:"adp_listEntities_outputJson,omitempty"`
	AdpCleanUpHistory                                      bool     `json:"adp_cleanUpHistory,omitempty"`
	AdpListEntitiesDescriptionSettingFilterValueDateFormat string   `json:"adp_listEntities_descriptionSettingFilterValueDateFormat,omitempty"`
	AdpListEntitiesDescriptionFilters                      []string `json:"adp_listEntities_descriptionFilters,omitempty"`
	AdpListEntitiesAxcServiceUser                          string   `json:"adp_listEntities_axcServiceUser,omitempty"`
	AdpListEntitiesAxcFields                               string   `json:"adp_listEntities_axcFields,omitempty"`
	AdpTaskTimeout                                         int      `json:"adp_taskTimeout,omitempty"`
	AdpListEntitiesHTTPSTrustCertificate                   string   `json:"adp_listEntities_httpsTrustCertificate,omitempty"`
	AdpListEntitiesHost                                    string   `json:"adp_listEntities_host,omitempty"`
	AdpListEntitiesOutputFilename                          string   `json:"adp_listEntities_outputFilename,omitempty"`
	AdpListEntitiesID                                      string   `json:"adp_listEntities_id,omitempty"`
	AdpListEntitiesHTTPSAllowUntrustedHosts                string   `json:"adp_listEntities_httpsAllowUntrustedHosts,omitempty"`
}

// EnableAdpLogging ...
func (c *ListEntitiesConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

// EnableAdpExecutionPersistent ...
func (c *ListEntitiesConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

// NewListEntitiesTaskRequest ...
func NewListEntitiesTaskRequest(opts ...func(*ListEntitiesConfiguration)) *TaskRequest {

	cfg := &ListEntitiesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &TaskRequest{
		TaskType:          "List Entities",
		TaskDescription:   "Writes a list of entities to an output variable",
		TaskConfiguration: cfg,
		TaskDisplayName:   "List Entities",
	}
}

// WithListEntitiesWhiteList ...
func WithListEntitiesWhiteList(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesWhiteList = s
	}
}

// WithListEntitiesRelatedEntity ...
func WithListEntitiesRelatedEntity(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesRelatedEntity = s
	}
}

// WithListEntitiesType ...
func WithListEntitiesType(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesType = s
	}
}

// WithListEntitiesID ...
func WithListEntitiesID(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesID = s
	}
}

// WithListEntitiesAxcFields ...
func WithListEntitiesAxcFields(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesAxcFields = s
	}
}

// WithListEntitiesStatus ...
func WithListEntitiesStatus(s string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesStatus = s
	}
}

// ListEntitiesExecutionMetaData ...
type ListEntitiesExecutionMetaData struct {
	AdpEntitiesOutputFileName string          `json:"adp_entities_output_file_name"`
	AdpEntitiesJSONOutput     json.RawMessage `json:"adp_entities_json_output"`
}

func (meta *ListEntitiesExecutionMetaData) Output() string {
	return string(meta.AdpEntitiesJSONOutput)
}
