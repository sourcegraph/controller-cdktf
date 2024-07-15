package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary struct {
	// canary_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#canary_deployment GoogleClouddeployDeliveryPipeline#canary_deployment}
	CanaryDeployment *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment `field:"optional" json:"canaryDeployment" yaml:"canaryDeployment"`
	// custom_canary_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#custom_canary_deployment GoogleClouddeployDeliveryPipeline#custom_canary_deployment}
	CustomCanaryDeployment *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment `field:"optional" json:"customCanaryDeployment" yaml:"customCanaryDeployment"`
	// runtime_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#runtime_config GoogleClouddeployDeliveryPipeline#runtime_config}
	RuntimeConfig *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
}

