package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshRouteDestinations struct {
	// Required.
	//
	// The clusters where the Gateway API HTTPRoute resource will be deployed to. Valid entries include the associated entities IDs configured in the Target resource and "@self" to include the Target cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_delivery_pipeline#destination_ids GoogleClouddeployDeliveryPipeline#destination_ids}
	DestinationIds *[]*string `field:"required" json:"destinationIds" yaml:"destinationIds"`
	// Optional.
	//
	// Whether to propagate the Kubernetes Service to the route destination clusters. The Service will always be deployed to the Target cluster even if the HTTPRoute is not. This option may be used to facilitiate successful DNS lookup in the route destination clusters. Can only be set to true if destinations are specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_delivery_pipeline#propagate_service GoogleClouddeployDeliveryPipeline#propagate_service}
	PropagateService interface{} `field:"optional" json:"propagateService" yaml:"propagateService"`
}

