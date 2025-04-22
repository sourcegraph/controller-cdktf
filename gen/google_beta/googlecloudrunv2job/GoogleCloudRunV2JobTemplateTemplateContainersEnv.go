package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_run_v2_job#name GoogleCloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Literal value of the environment variable.
	//
	// Defaults to "" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_run_v2_job#value GoogleCloudRunV2Job#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloud_run_v2_job#value_source GoogleCloudRunV2Job#value_source}
	ValueSource *GoogleCloudRunV2JobTemplateTemplateContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

