package googlestorageinsightsdatasetconfig


type GoogleStorageInsightsDatasetConfigSourceProjects struct {
	// The list of project numbers to include in the DatasetConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config#project_numbers GoogleStorageInsightsDatasetConfig#project_numbers}
	ProjectNumbers *[]*string `field:"optional" json:"projectNumbers" yaml:"projectNumbers"`
}

