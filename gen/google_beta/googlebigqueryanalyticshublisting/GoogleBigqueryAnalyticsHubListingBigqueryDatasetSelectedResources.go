package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingBigqueryDatasetSelectedResources struct {
	// Format: For routine: projects/{projectId}/datasets/{datasetId}/routines/{routineId} Example:"projects/test_project/datasets/test_dataset/routines/test_routine".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_listing#routine GoogleBigqueryAnalyticsHubListing#routine}
	Routine *string `field:"optional" json:"routine" yaml:"routine"`
	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_listing#table GoogleBigqueryAnalyticsHubListing#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
}

