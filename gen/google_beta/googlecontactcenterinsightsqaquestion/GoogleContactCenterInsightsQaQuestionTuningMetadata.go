package googlecontactcenterinsightsqaquestion


type GoogleContactCenterInsightsQaQuestionTuningMetadata struct {
	// A list of any applicable data validation warnings about the question's feedback labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#dataset_validation_warnings GoogleContactCenterInsightsQaQuestion#dataset_validation_warnings}
	DatasetValidationWarnings *[]*string `field:"optional" json:"datasetValidationWarnings" yaml:"datasetValidationWarnings"`
	// Total number of valid labels provided for the question at the time of tuining.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#total_valid_label_count GoogleContactCenterInsightsQaQuestion#total_valid_label_count}
	TotalValidLabelCount *string `field:"optional" json:"totalValidLabelCount" yaml:"totalValidLabelCount"`
	// Error status of the tuning operation for the question. Will only be set if the tuning operation failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question#tuning_error GoogleContactCenterInsightsQaQuestion#tuning_error}
	TuningError *string `field:"optional" json:"tuningError" yaml:"tuningError"`
}

