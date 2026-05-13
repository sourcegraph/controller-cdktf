package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport struct {
	// recipients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#recipients GoogleDataplexDatascan#recipients}
	Recipients *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients `field:"required" json:"recipients" yaml:"recipients"`
	// job_end_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#job_end_trigger GoogleDataplexDatascan#job_end_trigger}
	JobEndTrigger *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger `field:"optional" json:"jobEndTrigger" yaml:"jobEndTrigger"`
	// job_failure_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#job_failure_trigger GoogleDataplexDatascan#job_failure_trigger}
	JobFailureTrigger *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger `field:"optional" json:"jobFailureTrigger" yaml:"jobFailureTrigger"`
	// score_threshold_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#score_threshold_trigger GoogleDataplexDatascan#score_threshold_trigger}
	ScoreThresholdTrigger *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger `field:"optional" json:"scoreThresholdTrigger" yaml:"scoreThresholdTrigger"`
}

