package googlebigquerydataset


type GoogleBigqueryDatasetAccessView struct {
	// The ID of the dataset containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset#dataset_id GoogleBigqueryDataset#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset#project_id GoogleBigqueryDataset#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ID of the table.
	//
	// The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset#table_id GoogleBigqueryDataset#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

