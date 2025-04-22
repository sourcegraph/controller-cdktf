package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig struct {
	// whistle_config_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#whistle_config_source GoogleHealthcarePipelineJob#whistle_config_source}
	WhistleConfigSource *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource `field:"required" json:"whistleConfigSource" yaml:"whistleConfigSource"`
	// Describes the mapping configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_healthcare_pipeline_job#description GoogleHealthcarePipelineJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

