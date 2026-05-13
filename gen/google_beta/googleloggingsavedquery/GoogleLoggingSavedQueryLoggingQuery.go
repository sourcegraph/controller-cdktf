package googleloggingsavedquery


type GoogleLoggingSavedQueryLoggingQuery struct {
	// An [advanced logs filter](https://cloud.google.com/logging/docs/view/advanced-filters) which is used to match log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#filter GoogleLoggingSavedQuery#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Characters will be counted from the end of the string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#summary_field_end GoogleLoggingSavedQuery#summary_field_end}
	SummaryFieldEnd *float64 `field:"optional" json:"summaryFieldEnd" yaml:"summaryFieldEnd"`
	// summary_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#summary_fields GoogleLoggingSavedQuery#summary_fields}
	SummaryFields interface{} `field:"optional" json:"summaryFields" yaml:"summaryFields"`
	// Characters will be counted from the start of the string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#summary_field_start GoogleLoggingSavedQuery#summary_field_start}
	SummaryFieldStart *float64 `field:"optional" json:"summaryFieldStart" yaml:"summaryFieldStart"`
}

