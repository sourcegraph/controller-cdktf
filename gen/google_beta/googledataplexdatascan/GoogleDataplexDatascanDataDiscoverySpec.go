package googledataplexdatascan


type GoogleDataplexDatascanDataDiscoverySpec struct {
	// bigquery_publishing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#bigquery_publishing_config GoogleDataplexDatascan#bigquery_publishing_config}
	BigqueryPublishingConfig *GoogleDataplexDatascanDataDiscoverySpecBigqueryPublishingConfig `field:"optional" json:"bigqueryPublishingConfig" yaml:"bigqueryPublishingConfig"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#storage_config GoogleDataplexDatascan#storage_config}
	StorageConfig *GoogleDataplexDatascanDataDiscoverySpecStorageConfig `field:"optional" json:"storageConfig" yaml:"storageConfig"`
}

