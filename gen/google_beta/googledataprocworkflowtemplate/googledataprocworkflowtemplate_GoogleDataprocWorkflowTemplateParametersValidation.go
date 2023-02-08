package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateParametersValidation struct {
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#regex GoogleDataprocWorkflowTemplate#regex}
	Regex *GoogleDataprocWorkflowTemplateParametersValidationRegex `field:"optional" json:"regex" yaml:"regex"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#values GoogleDataprocWorkflowTemplate#values}
	Values *GoogleDataprocWorkflowTemplateParametersValidationValues `field:"optional" json:"values" yaml:"values"`
}

