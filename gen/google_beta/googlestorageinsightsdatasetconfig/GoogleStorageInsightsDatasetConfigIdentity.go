package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigIdentity struct {
	// Type of identity to use for the DatasetConfig. Possible values: ["IDENTITY_TYPE_PER_CONFIG", "IDENTITY_TYPE_PER_PROJECT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config#type GoogleStorageInsightsDatasetConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

