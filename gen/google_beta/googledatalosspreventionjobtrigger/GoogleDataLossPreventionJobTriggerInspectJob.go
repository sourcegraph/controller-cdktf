package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJob struct {
	// The name of the template to run when this job is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_loss_prevention_job_trigger#inspect_template_name GoogleDataLossPreventionJobTrigger#inspect_template_name}
	InspectTemplateName *string `field:"required" json:"inspectTemplateName" yaml:"inspectTemplateName"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_loss_prevention_job_trigger#storage_config GoogleDataLossPreventionJobTrigger#storage_config}
	StorageConfig *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig `field:"required" json:"storageConfig" yaml:"storageConfig"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_loss_prevention_job_trigger#actions GoogleDataLossPreventionJobTrigger#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// inspect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_loss_prevention_job_trigger#inspect_config GoogleDataLossPreventionJobTrigger#inspect_config}
	InspectConfig *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig `field:"optional" json:"inspectConfig" yaml:"inspectConfig"`
}

