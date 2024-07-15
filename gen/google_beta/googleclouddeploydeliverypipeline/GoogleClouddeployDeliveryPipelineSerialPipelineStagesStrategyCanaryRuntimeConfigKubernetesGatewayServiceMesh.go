package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh struct {
	// Required. Name of the Kubernetes Deployment whose traffic is managed by the specified HTTPRoute and Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#deployment GoogleClouddeployDeliveryPipeline#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// Required. Name of the Gateway API HTTPRoute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#http_route GoogleClouddeployDeliveryPipeline#http_route}
	HttpRoute *string `field:"required" json:"httpRoute" yaml:"httpRoute"`
	// Required. Name of the Kubernetes Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#service GoogleClouddeployDeliveryPipeline#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional.
	//
	// The time to wait for route updates to propagate. The maximum configurable time is 3 hours, in seconds format. If unspecified, there is no wait time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#route_update_wait_time GoogleClouddeployDeliveryPipeline#route_update_wait_time}
	RouteUpdateWaitTime *string `field:"optional" json:"routeUpdateWaitTime" yaml:"routeUpdateWaitTime"`
	// Optional.
	//
	// The amount of time to migrate traffic back from the canary Service to the original Service during the stable phase deployment. If specified, must be between 15s and 3600s. If unspecified, there is no cutback time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#stable_cutback_duration GoogleClouddeployDeliveryPipeline#stable_cutback_duration}
	StableCutbackDuration *string `field:"optional" json:"stableCutbackDuration" yaml:"stableCutbackDuration"`
}

