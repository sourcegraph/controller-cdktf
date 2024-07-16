package googlebigqueryanalyticshublistingiambinding


type GoogleBigqueryAnalyticsHubListingIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_analytics_hub_listing_iam_binding#expression GoogleBigqueryAnalyticsHubListingIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_analytics_hub_listing_iam_binding#title GoogleBigqueryAnalyticsHubListingIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_analytics_hub_listing_iam_binding#description GoogleBigqueryAnalyticsHubListingIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

