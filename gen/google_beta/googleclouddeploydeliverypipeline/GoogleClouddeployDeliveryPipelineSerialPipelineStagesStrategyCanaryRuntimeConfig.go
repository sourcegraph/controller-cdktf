package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig struct {
	// cloud_run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#cloud_run GoogleClouddeployDeliveryPipeline#cloud_run}
	CloudRun *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun `field:"optional" json:"cloudRun" yaml:"cloudRun"`
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#kubernetes GoogleClouddeployDeliveryPipeline#kubernetes}
	Kubernetes *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
}

