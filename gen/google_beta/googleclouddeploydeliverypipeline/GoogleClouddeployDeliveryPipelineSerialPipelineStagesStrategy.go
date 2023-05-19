package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategy struct {
	// canary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_clouddeploy_delivery_pipeline#canary GoogleClouddeployDeliveryPipeline#canary}
	Canary *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary `field:"optional" json:"canary" yaml:"canary"`
	// standard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_clouddeploy_delivery_pipeline#standard GoogleClouddeployDeliveryPipeline#standard}
	Standard *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard `field:"optional" json:"standard" yaml:"standard"`
}

