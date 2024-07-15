package googlebigquerydatasetaccess


type GoogleBigqueryDatasetAccessDatasetDatasetA struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#dataset_id GoogleBigqueryDatasetAccessA#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#project_id GoogleBigqueryDatasetAccessA#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

