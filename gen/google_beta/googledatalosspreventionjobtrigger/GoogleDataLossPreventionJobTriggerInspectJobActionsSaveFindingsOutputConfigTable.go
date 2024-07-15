package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTable struct {
	// Dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#dataset_id GoogleDataLossPreventionJobTrigger#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The Google Cloud Platform project ID of the project containing the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#project_id GoogleDataLossPreventionJobTrigger#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Name of the table.
	//
	// If is not set a new one will be generated for you with the following format:
	// 'dlp_googleapis_yyyy_mm_dd_[dlp_job_id]'. Pacific timezone will be used for generating the date details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#table_id GoogleDataLossPreventionJobTrigger#table_id}
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
}

