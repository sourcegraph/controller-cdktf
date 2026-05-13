package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingPubsubTopic struct {
	// Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_listing#topic GoogleBigqueryAnalyticsHubListing#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Region hint on where the data might be published.
	//
	// Data affinity regions are modifiable.
	// See https://cloud.google.com/about/locations for full listing of possible Cloud regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_analytics_hub_listing#data_affinity_regions GoogleBigqueryAnalyticsHubListing#data_affinity_regions}
	DataAffinityRegions *[]*string `field:"optional" json:"dataAffinityRegions" yaml:"dataAffinityRegions"`
}

