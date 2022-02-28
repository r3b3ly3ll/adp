package task

import (
	"encoding/json"
)

type TaxonomyStatisticConfiguration struct {
	AdpProgressTaskTimeout                                int                 `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaxonomyStatisticOutputJSONAbsFilePath             string              `json:"adp_taxonomyStatistic_outputJsonAbsFilePath,omitempty"`
	AdpTaxonomyStatisticApplicationIdentifier             string              `json:"adp_taxonomyStatistic_applicationIdentifier"`
	AdpTaskActive                                         bool                `json:"adp_taskActive,omitempty"`
	AdpTaxonomyStatisticAdpTaxonomyStatisticMainQueryType interface{}         `json:"adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType,omitempty"`
	AdpExecutionPersistent                                bool                `json:"adp_executionPersistent,omitempty"`
	AdpTaxonomyStatisticEngineUserName                    string              `json:"adp_taxonomyStatistic_engineUserName,omitempty"`
	AdpAbortWfOnFailure                                   bool                `json:"adp_abortWfOnFailure,omitempty"`
	AdpTaxonomyStatisticApplicationType                   string              `json:"adp_taxonomyStatistic_applicationType,omitempty"`
	AdpTaxonomyStatisticComputeCounts                     string              `json:"adp_taxonomyStatistic_computeCounts,omitempty"`
	AdpLoggingEnabled                                     bool                `json:"adp_loggingEnabled,omitempty"`
	AdpTaxonomyStatisticOutputJSONFilePath                string              `json:"adp_taxonomyStatistic_outputJsonFilePath,omitempty"`
	AdpTaxonomyStatisticEngineTaxonomies                  []EngineTaxonomyArg `json:"adp_taxonomyStatistic_engineTaxonomies"`
	AdpTaxonomyStatisticEngineUserPassword                string              `json:"adp_taxonomyStatistic_engineUserPassword,omitempty"`
	AdpTaxonomyStatisticOutputXMLAbsFilePath              string              `json:"adp_taxonomyStatistic_outputXmlAbsFilePath,omitempty"`
	AdpTaxonomyStatisticEngineQuery                       string              `json:"adp_taxonomyStatistic_engineQuery,omitempty"`
	AdpTaxonomyStatisticListCategoryProperties            string              `json:"adp_taxonomyStatistic_listCategoryProperties,omitempty"`
	AdpTaxonomyStatisticOutputTaxonomies                  []interface{}       `json:"adp_taxonomyStatistic_outputTaxonomies,omitempty"`
	AdpTaxonomyStatisticOutputJSON                        string              `json:"adp_taxonomyStatistic_outputJson,omitempty"`
	AdpTaxonomyStatisticEngineType                        string              `json:"adp_taxonomyStatistic_engineType,omitempty"`
	AdpCleanUpHistory                                     bool                `json:"adp_cleanUpHistory,omitempty"`
	AdpTaxonomyStatisticOutputXMLFilePath                 string              `json:"adp_taxonomyStatistic_outputXmlFilePath,omitempty"`
	AdpTaxonomyStatisticOutputFields                      []interface{}       `json:"adp_taxonomyStatistic_outputFields,omitempty"`
	AdpTaxonomyStatisticEngineGlobalSearch                string              `json:"adp_taxonomyStatistic_engineGlobalSearch,omitempty"`
	AdpTaxonomyStatisticListDocuments                     string              `json:"adp_taxonomyStatistic_listDocuments,omitempty"`
	AdpTaskTimeout                                        int                 `json:"adp_taskTimeout,omitempty"`
	AdpTaxonomyStatisticEngineName                        string              `json:"adp_taxonomyStatistic_engineName"`
}

func (c *TaxonomyStatisticConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

func (c *TaxonomyStatisticConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

func NewTaxonomyStatisticTaskRequest(opts ...func(*TaxonomyStatisticConfiguration)) *Request {
	cfg := &TaxonomyStatisticConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Taxonomy Statistic",
		TaskDescription:   "Retrieves category counts for a taxonomy",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Taxonomy statistic",
	}
}

// @TaskModelParameter(name="adp_taxonomyStatistic_engineTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Negation|Query", columnTypes="String|boolean|String", separator="|")
func WithTaxonomyStatisticEngineTaxonomies(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		engineTaxonomies := parseEngineTaxonomiesArgs(s)
		if len(engineTaxonomies) == 0 {
			return
		}
		c.AdpTaxonomyStatisticEngineTaxonomies = engineTaxonomies
	}
}

// ApplicationIdentifier and EngineName are mutually exclusive.
// If one is assigned, the other should be empty and can not be omitted.
func WithTaxonomyStatisticApplicationIdentifier(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(s) > 0 {
			c.AdpTaxonomyStatisticApplicationIdentifier = s
		}
	}
}

func WithTaxonomyStatisticEngineName(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(s) > 0 {
			c.AdpTaxonomyStatisticEngineName = s
		}
	}
}

func WithTaxonomyStatisticComputeCounts(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticComputeCounts = s
	}
}

func WithTaxonomyStatisticListCategoryProperties(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticListCategoryProperties = s
	}
}

func WithTaxonomyStatisticEngineUserName(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserName = s
	}
}

func WithTaxonomyStatisticEngineUserPassword(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserPassword = s
	}
}

type OutputTaxonomies struct {
	Taxonomy                  string `json:"Taxonomy"`
	Mode                      string `json:"Mode"`
	MaximumNumberOfCategories int    `json:"Maximum number of categories"`
}

// @TaskModelParameter(name="adp_taxonomyStatistic_outputTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Mode|Maximum number of categories", columnTypes="String|String|integer", separator="|")
func WithTaxonomyStatisticOutputTaxonomies(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		ot := OutputTaxonomies{
			Taxonomy:                  s,
			Mode:                      "Category counts",
			MaximumNumberOfCategories: 1000,
		}

		c.AdpTaxonomyStatisticOutputTaxonomies = append(c.AdpTaxonomyStatisticOutputTaxonomies, ot)
	}
}

type TaxonomyStatisticExecutionMetaData struct {
	AdpTaxonomyStatisticsJsonOutput json.RawMessage `json:"adp_taxonomy_statistics_json_output"`
}

func (meta *TaxonomyStatisticExecutionMetaData) Output() string {

	output := string(meta.AdpTaxonomyStatisticsJsonOutput)
	unquoteJSONOutput(&output)

	return output
}

func NewTaxonomyStatisticTaskResponse() *Response {
	return &Response{
		ExecutionMetaData: &TaxonomyStatisticExecutionMetaData{},
	}
}
