package googlebigqueryanalyticshublistingsubscription


type GoogleBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReference struct {
	// A unique ID for this dataset, without the project name.
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#dataset_id GoogleBigqueryAnalyticsHubListingSubscription#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_analytics_hub_listing_subscription#project_id GoogleBigqueryAnalyticsHubListingSubscription#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

