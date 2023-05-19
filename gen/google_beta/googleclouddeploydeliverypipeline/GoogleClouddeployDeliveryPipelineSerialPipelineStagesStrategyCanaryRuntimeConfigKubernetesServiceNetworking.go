package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking struct {
	// Required. Name of the Kubernetes Deployment whose traffic is managed by the specified Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_clouddeploy_delivery_pipeline#deployment GoogleClouddeployDeliveryPipeline#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// Required. Name of the Kubernetes Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_clouddeploy_delivery_pipeline#service GoogleClouddeployDeliveryPipeline#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

