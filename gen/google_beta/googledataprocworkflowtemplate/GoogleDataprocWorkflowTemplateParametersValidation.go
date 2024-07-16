package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateParametersValidation struct {
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#regex GoogleDataprocWorkflowTemplate#regex}
	Regex *GoogleDataprocWorkflowTemplateParametersValidationRegex `field:"optional" json:"regex" yaml:"regex"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#values GoogleDataprocWorkflowTemplate#values}
	Values *GoogleDataprocWorkflowTemplateParametersValidationValues `field:"optional" json:"values" yaml:"values"`
}

