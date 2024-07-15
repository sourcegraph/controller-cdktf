package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig struct {
	// If this template is specified, it will serve as the default de-identify template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#deidentify_template GoogleDataLossPreventionJobTrigger#deidentify_template}
	DeidentifyTemplate *string `field:"optional" json:"deidentifyTemplate" yaml:"deidentifyTemplate"`
	// If this template is specified, it will serve as the de-identify template for images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#image_redact_template GoogleDataLossPreventionJobTrigger#image_redact_template}
	ImageRedactTemplate *string `field:"optional" json:"imageRedactTemplate" yaml:"imageRedactTemplate"`
	// If this template is specified, it will serve as the de-identify template for structured content such as delimited files and tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#structured_deidentify_template GoogleDataLossPreventionJobTrigger#structured_deidentify_template}
	StructuredDeidentifyTemplate *string `field:"optional" json:"structuredDeidentifyTemplate" yaml:"structuredDeidentifyTemplate"`
}

