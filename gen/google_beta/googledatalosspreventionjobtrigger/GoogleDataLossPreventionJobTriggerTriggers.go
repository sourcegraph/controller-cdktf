package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerTriggers struct {
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#manual GoogleDataLossPreventionJobTrigger#manual}
	Manual *GoogleDataLossPreventionJobTriggerTriggersManual `field:"optional" json:"manual" yaml:"manual"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#schedule GoogleDataLossPreventionJobTrigger#schedule}
	Schedule *GoogleDataLossPreventionJobTriggerTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

