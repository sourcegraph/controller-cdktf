package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingBigqueryDataset struct {
	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_analytics_hub_listing#dataset GoogleBigqueryAnalyticsHubListing#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// selected_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_analytics_hub_listing#selected_resources GoogleBigqueryAnalyticsHubListing#selected_resources}
	SelectedResources interface{} `field:"optional" json:"selectedResources" yaml:"selectedResources"`
}

