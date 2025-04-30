package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and may not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Literal value of the environment variable.
	//
	// Defaults to "" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#value GoogleCloudRunV2Service#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_run_v2_service#value_source GoogleCloudRunV2Service#value_source}
	ValueSource *GoogleCloudRunV2ServiceTemplateContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

