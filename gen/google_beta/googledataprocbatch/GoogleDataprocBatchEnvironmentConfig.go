package googledataprocbatch


type GoogleDataprocBatchEnvironmentConfig struct {
	// execution_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#execution_config GoogleDataprocBatch#execution_config}
	ExecutionConfig *GoogleDataprocBatchEnvironmentConfigExecutionConfig `field:"optional" json:"executionConfig" yaml:"executionConfig"`
	// peripherals_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#peripherals_config GoogleDataprocBatch#peripherals_config}
	PeripheralsConfig *GoogleDataprocBatchEnvironmentConfigPeripheralsConfig `field:"optional" json:"peripheralsConfig" yaml:"peripheralsConfig"`
}

