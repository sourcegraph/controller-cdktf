package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations struct {
	// The list of cloud storage locations to exclude in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config#locations GoogleStorageInsightsDatasetConfig#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
}

