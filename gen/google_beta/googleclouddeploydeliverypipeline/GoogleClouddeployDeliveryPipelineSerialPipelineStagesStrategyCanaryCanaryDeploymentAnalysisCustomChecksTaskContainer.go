package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentAnalysisCustomChecksTaskContainer struct {
	// Required. Image is the container image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#image GoogleClouddeployDeliveryPipeline#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Optional. Args is the container arguments to use. This overrides the default arguments defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#args GoogleClouddeployDeliveryPipeline#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Optional. Command is the container entrypoint to use. This overrides the default entrypoint defined in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#command GoogleClouddeployDeliveryPipeline#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Optional. Environment variables that are set in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_clouddeploy_delivery_pipeline#env GoogleClouddeployDeliveryPipeline#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

