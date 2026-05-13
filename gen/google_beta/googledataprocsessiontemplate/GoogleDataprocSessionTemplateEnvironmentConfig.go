package googledataprocsessiontemplate


type GoogleDataprocSessionTemplateEnvironmentConfig struct {
	// execution_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#execution_config GoogleDataprocSessionTemplate#execution_config}
	ExecutionConfig *GoogleDataprocSessionTemplateEnvironmentConfigExecutionConfig `field:"optional" json:"executionConfig" yaml:"executionConfig"`
	// peripherals_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_session_template#peripherals_config GoogleDataprocSessionTemplate#peripherals_config}
	PeripheralsConfig *GoogleDataprocSessionTemplateEnvironmentConfigPeripheralsConfig `field:"optional" json:"peripheralsConfig" yaml:"peripheralsConfig"`
}

