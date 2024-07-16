package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub struct {
	// Cloud Pub/Sub topic to send notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#topic GoogleDataLossPreventionJobTrigger#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

