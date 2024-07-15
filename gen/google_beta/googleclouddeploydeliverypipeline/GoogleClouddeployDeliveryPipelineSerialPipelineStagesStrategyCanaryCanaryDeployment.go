package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment struct {
	// Required.
	//
	// The percentage based deployments that will occur as a part of a `Rollout`. List is expected in ascending order and each integer n is 0 <= n < 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#percentages GoogleClouddeployDeliveryPipeline#percentages}
	Percentages *[]*float64 `field:"required" json:"percentages" yaml:"percentages"`
	// postdeploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#postdeploy GoogleClouddeployDeliveryPipeline#postdeploy}
	Postdeploy *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy `field:"optional" json:"postdeploy" yaml:"postdeploy"`
	// predeploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#predeploy GoogleClouddeployDeliveryPipeline#predeploy}
	Predeploy *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy `field:"optional" json:"predeploy" yaml:"predeploy"`
	// Whether to run verify tests after each percentage deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#verify GoogleClouddeployDeliveryPipeline#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
}

