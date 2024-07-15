package googlebigqueryjob


type GoogleBigqueryJobExtractSourceModel struct {
	// The ID of the dataset containing this model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#dataset_id GoogleBigqueryJob#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#model_id GoogleBigqueryJob#model_id}
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// The ID of the project containing this model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#project_id GoogleBigqueryJob#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

