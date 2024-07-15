package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard struct {
	// postdeploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#postdeploy GoogleClouddeployDeliveryPipeline#postdeploy}
	Postdeploy *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy `field:"optional" json:"postdeploy" yaml:"postdeploy"`
	// predeploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#predeploy GoogleClouddeployDeliveryPipeline#predeploy}
	Predeploy *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy `field:"optional" json:"predeploy" yaml:"predeploy"`
	// Whether to verify a deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#verify GoogleClouddeployDeliveryPipeline#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
}

