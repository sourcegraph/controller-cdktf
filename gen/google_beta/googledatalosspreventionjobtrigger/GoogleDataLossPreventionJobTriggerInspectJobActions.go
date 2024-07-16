package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActions struct {
	// deidentify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#deidentify GoogleDataLossPreventionJobTrigger#deidentify}
	Deidentify *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify `field:"optional" json:"deidentify" yaml:"deidentify"`
	// job_notification_emails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#job_notification_emails GoogleDataLossPreventionJobTrigger#job_notification_emails}
	JobNotificationEmails *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails `field:"optional" json:"jobNotificationEmails" yaml:"jobNotificationEmails"`
	// publish_findings_to_cloud_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#publish_findings_to_cloud_data_catalog GoogleDataLossPreventionJobTrigger#publish_findings_to_cloud_data_catalog}
	PublishFindingsToCloudDataCatalog *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog `field:"optional" json:"publishFindingsToCloudDataCatalog" yaml:"publishFindingsToCloudDataCatalog"`
	// publish_summary_to_cscc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#publish_summary_to_cscc GoogleDataLossPreventionJobTrigger#publish_summary_to_cscc}
	PublishSummaryToCscc *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc `field:"optional" json:"publishSummaryToCscc" yaml:"publishSummaryToCscc"`
	// publish_to_stackdriver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#publish_to_stackdriver GoogleDataLossPreventionJobTrigger#publish_to_stackdriver}
	PublishToStackdriver *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver `field:"optional" json:"publishToStackdriver" yaml:"publishToStackdriver"`
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#pub_sub GoogleDataLossPreventionJobTrigger#pub_sub}
	PubSub *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// save_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#save_findings GoogleDataLossPreventionJobTrigger#save_findings}
	SaveFindings *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings `field:"optional" json:"saveFindings" yaml:"saveFindings"`
}

