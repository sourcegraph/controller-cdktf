package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJob struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#actions GoogleDataLossPreventionJobTrigger#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the template to run when this job is triggered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#inspect_template_name GoogleDataLossPreventionJobTrigger#inspect_template_name}
	InspectTemplateName *string `field:"required" json:"inspectTemplateName" yaml:"inspectTemplateName"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#storage_config GoogleDataLossPreventionJobTrigger#storage_config}
	StorageConfig *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig `field:"required" json:"storageConfig" yaml:"storageConfig"`
}

