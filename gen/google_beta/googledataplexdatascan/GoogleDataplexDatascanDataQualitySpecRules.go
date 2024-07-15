package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRules struct {
	// The dimension a rule belongs to.
	//
	// Results are also aggregated at the dimension level. Supported dimensions are ["COMPLETENESS", "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS", "INTEGRITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#dimension GoogleDataplexDatascan#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// The unnested column which this rule is evaluated against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#column GoogleDataplexDatascan#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// Description of the rule. The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#description GoogleDataplexDatascan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Rows with null values will automatically fail a rule, unless ignoreNull is true.
	//
	// In that case, such null rows are trivially considered passing. Only applicable to ColumnMap rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#ignore_null GoogleDataplexDatascan#ignore_null}
	IgnoreNull interface{} `field:"optional" json:"ignoreNull" yaml:"ignoreNull"`
	// A mutable name for the rule.
	//
	// The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
	// The maximum length is 63 characters.
	// Must start with a letter.
	// Must end with a number or a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#name GoogleDataplexDatascan#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// non_null_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#non_null_expectation GoogleDataplexDatascan#non_null_expectation}
	NonNullExpectation *GoogleDataplexDatascanDataQualitySpecRulesNonNullExpectation `field:"optional" json:"nonNullExpectation" yaml:"nonNullExpectation"`
	// range_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#range_expectation GoogleDataplexDatascan#range_expectation}
	RangeExpectation *GoogleDataplexDatascanDataQualitySpecRulesRangeExpectation `field:"optional" json:"rangeExpectation" yaml:"rangeExpectation"`
	// regex_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#regex_expectation GoogleDataplexDatascan#regex_expectation}
	RegexExpectation *GoogleDataplexDatascanDataQualitySpecRulesRegexExpectation `field:"optional" json:"regexExpectation" yaml:"regexExpectation"`
	// row_condition_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#row_condition_expectation GoogleDataplexDatascan#row_condition_expectation}
	RowConditionExpectation *GoogleDataplexDatascanDataQualitySpecRulesRowConditionExpectation `field:"optional" json:"rowConditionExpectation" yaml:"rowConditionExpectation"`
	// set_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#set_expectation GoogleDataplexDatascan#set_expectation}
	SetExpectation *GoogleDataplexDatascanDataQualitySpecRulesSetExpectation `field:"optional" json:"setExpectation" yaml:"setExpectation"`
	// sql_assertion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#sql_assertion GoogleDataplexDatascan#sql_assertion}
	SqlAssertion *GoogleDataplexDatascanDataQualitySpecRulesSqlAssertion `field:"optional" json:"sqlAssertion" yaml:"sqlAssertion"`
	// statistic_range_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#statistic_range_expectation GoogleDataplexDatascan#statistic_range_expectation}
	StatisticRangeExpectation *GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation `field:"optional" json:"statisticRangeExpectation" yaml:"statisticRangeExpectation"`
	// table_condition_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#table_condition_expectation GoogleDataplexDatascan#table_condition_expectation}
	TableConditionExpectation *GoogleDataplexDatascanDataQualitySpecRulesTableConditionExpectation `field:"optional" json:"tableConditionExpectation" yaml:"tableConditionExpectation"`
	// The minimum ratio of passing_rows / total_rows required to pass this rule, with a range of [0.0, 1.0]. 0 indicates default value (i.e. 1.0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#threshold GoogleDataplexDatascan#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// uniqueness_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_datascan#uniqueness_expectation GoogleDataplexDatascan#uniqueness_expectation}
	UniquenessExpectation *GoogleDataplexDatascanDataQualitySpecRulesUniquenessExpectation `field:"optional" json:"uniquenessExpectation" yaml:"uniquenessExpectation"`
}

