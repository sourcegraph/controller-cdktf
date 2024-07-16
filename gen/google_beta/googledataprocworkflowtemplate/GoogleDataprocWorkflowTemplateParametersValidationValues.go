package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateParametersValidationValues struct {
	// Required. List of allowed values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#values GoogleDataprocWorkflowTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

