package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingBigqueryDatasetSelectedResources struct {
	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_analytics_hub_listing#table GoogleBigqueryAnalyticsHubListing#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
}

