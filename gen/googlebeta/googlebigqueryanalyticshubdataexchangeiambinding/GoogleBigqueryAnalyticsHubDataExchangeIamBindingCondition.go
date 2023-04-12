package googlebigqueryanalyticshubdataexchangeiambinding


type GoogleBigqueryAnalyticsHubDataExchangeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_analytics_hub_data_exchange_iam_binding#expression GoogleBigqueryAnalyticsHubDataExchangeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_analytics_hub_data_exchange_iam_binding#title GoogleBigqueryAnalyticsHubDataExchangeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_analytics_hub_data_exchange_iam_binding#description GoogleBigqueryAnalyticsHubDataExchangeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

