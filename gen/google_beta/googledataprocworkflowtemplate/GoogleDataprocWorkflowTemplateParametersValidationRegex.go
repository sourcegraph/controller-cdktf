package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateParametersValidationRegex struct {
	// Required.
	//
	// RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#regexes GoogleDataprocWorkflowTemplate#regexes}
	Regexes *[]*string `field:"required" json:"regexes" yaml:"regexes"`
}

