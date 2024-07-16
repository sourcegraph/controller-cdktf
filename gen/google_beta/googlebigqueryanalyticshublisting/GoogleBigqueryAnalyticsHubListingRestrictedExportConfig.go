package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingRestrictedExportConfig struct {
	// If true, enable restricted export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_analytics_hub_listing#enabled GoogleBigqueryAnalyticsHubListing#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If true, restrict export of query result derived from restricted linked dataset table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_analytics_hub_listing#restrict_query_result GoogleBigqueryAnalyticsHubListing#restrict_query_result}
	RestrictQueryResult interface{} `field:"optional" json:"restrictQueryResult" yaml:"restrictQueryResult"`
}

