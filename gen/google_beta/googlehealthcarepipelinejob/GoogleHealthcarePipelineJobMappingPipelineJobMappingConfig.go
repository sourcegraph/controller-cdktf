package googlehealthcarepipelinejob


type GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig struct {
	// Describes the mapping configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_healthcare_pipeline_job#description GoogleHealthcarePipelineJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// whistle_config_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_healthcare_pipeline_job#whistle_config_source GoogleHealthcarePipelineJob#whistle_config_source}
	WhistleConfigSource *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource `field:"optional" json:"whistleConfigSource" yaml:"whistleConfigSource"`
}

