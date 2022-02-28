package task

import "strings"

// Delimiters used for parsing Args.
const (
	ARGDELIMITER = ";"
	NE           = "!="
	EQ           = "="
)

// AdvancedRestrictionsArg ...
// Fuzzy Search is currently set to false.
type AdvancedRestrictionsArg struct {
	AdvancedSearchField string `json:"Advanced Search Field"`
	Negation            bool   `json:"Negation"`
	FuzzySearch         bool   `json:"Fuzzy Search"`
	AdvancedExpression  string `json:"Advanced Expression"`
}

func parseAdvancedRestrictionsArg(s string) []AdvancedRestrictionsArg {
	var advancedRestrictionsArgs []AdvancedRestrictionsArg

	for _, restriction := range strings.Split(s, ARGDELIMITER) {
		var fields []string
		var negation bool

		if strings.Contains(restriction, NE) {
			negation = true
			fields = strings.SplitN(restriction, NE, 2)
		} else {
			fields = strings.SplitN(restriction, EQ, 2)
		}

		if len(fields) == 2 {
			fld := strings.TrimSpace(fields[0])
			exp := strings.TrimSpace(fields[1])

			if len(fld) > 0 {
				advancedRestrictionsArgs = append(advancedRestrictionsArgs,
					AdvancedRestrictionsArg{
						AdvancedSearchField: fld,
						Negation:            negation,
						FuzzySearch:         false,
						AdvancedExpression:  exp,
					})
			}
		}
	}
	return advancedRestrictionsArgs
}

// EngineTaxonomyArg represents data structure for category/taxonmy query statement,
// e.g., csv_guts_datatype=email.
type EngineTaxonomyArg struct {
	Taxonomy string `json:"Taxonomy"`
	Negation bool   `json:"Negation"`
	Query    string `json:"Query"`
}

func parseEngineTaxonomiesArgs(s string) []EngineTaxonomyArg {
	var engineTaxonomyArgs []EngineTaxonomyArg

	for _, taxonomy := range strings.Split(s, ARGDELIMITER) {
		var fields []string
		var negation bool

		if strings.Contains(taxonomy, NE) {
			negation = true
			fields = strings.SplitN(taxonomy, NE, 2)
		} else {
			fields = strings.SplitN(taxonomy, EQ, 2)
		}

		if len(fields) == 2 {
			// remove leading and trailing spaces of each field
			tax := strings.TrimSpace(fields[0])
			qry := strings.TrimSpace(fields[1])

			if len(tax) > 0 {
				engineTaxonomyArgs = append(engineTaxonomyArgs,
					EngineTaxonomyArg{
						Taxonomy: tax,
						Negation: negation,
						Query:    qry,
					})
			}
		}
	}

	return engineTaxonomyArgs
}
