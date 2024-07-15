package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference struct {
	// The dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#dataset_id GoogleDataLossPreventionJobTrigger#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The Google Cloud Platform project ID of the project containing the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#project_id GoogleDataLossPreventionJobTrigger#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#table_id GoogleDataLossPreventionJobTrigger#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

