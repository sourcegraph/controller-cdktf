package googlebigqueryjob


type GoogleBigqueryJobQueryDefaultDataset struct {
	// The dataset. Can be specified '{{dataset_id}}' if 'project_id' is also set, or of the form 'projects/{{project}}/datasets/{{dataset_id}}' if not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#dataset_id GoogleBigqueryJob#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The ID of the project containing this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#project_id GoogleBigqueryJob#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

