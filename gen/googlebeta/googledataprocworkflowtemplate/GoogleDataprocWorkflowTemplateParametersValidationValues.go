package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateParametersValidationValues struct {
	// Required. List of allowed values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#values GoogleDataprocWorkflowTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

