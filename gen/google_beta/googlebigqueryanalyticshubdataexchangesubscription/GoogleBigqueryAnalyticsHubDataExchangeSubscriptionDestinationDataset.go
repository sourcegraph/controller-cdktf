package googlebigqueryanalyticshubdataexchangesubscription


type GoogleBigqueryAnalyticsHubDataExchangeSubscriptionDestinationDataset struct {
	// dataset_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_data_exchange_subscription#dataset_reference GoogleBigqueryAnalyticsHubDataExchangeSubscription#dataset_reference}
	DatasetReference *GoogleBigqueryAnalyticsHubDataExchangeSubscriptionDestinationDatasetDatasetReference `field:"required" json:"datasetReference" yaml:"datasetReference"`
	// The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_data_exchange_subscription#location GoogleBigqueryAnalyticsHubDataExchangeSubscription#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A user-friendly description of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_data_exchange_subscription#description GoogleBigqueryAnalyticsHubDataExchangeSubscription#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A descriptive name for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_data_exchange_subscription#friendly_name GoogleBigqueryAnalyticsHubDataExchangeSubscription#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// The labels associated with this dataset. You can use these to organize and group your datasets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_data_exchange_subscription#labels GoogleBigqueryAnalyticsHubDataExchangeSubscription#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

