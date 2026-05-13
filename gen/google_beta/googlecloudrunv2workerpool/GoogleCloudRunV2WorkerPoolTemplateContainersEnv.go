package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and may not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#name GoogleCloudRunV2WorkerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Literal value of the environment variable.
	//
	// Defaults to "" and the maximum allowed length is 32768 characters. Variable references are not supported in Cloud Run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#value GoogleCloudRunV2WorkerPool#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_v2_worker_pool#value_source GoogleCloudRunV2WorkerPool#value_source}
	ValueSource *GoogleCloudRunV2WorkerPoolTemplateContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

