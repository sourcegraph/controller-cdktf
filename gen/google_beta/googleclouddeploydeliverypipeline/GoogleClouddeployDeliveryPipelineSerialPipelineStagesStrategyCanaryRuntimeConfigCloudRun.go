package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun struct {
	// Whether Cloud Deploy should update the traffic stanza in a Cloud Run Service on the user's behalf to facilitate traffic splitting.
	//
	// This is required to be true for CanaryDeployments, but optional for CustomCanaryDeployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#automatic_traffic_control GoogleClouddeployDeliveryPipeline#automatic_traffic_control}
	AutomaticTrafficControl interface{} `field:"optional" json:"automaticTrafficControl" yaml:"automaticTrafficControl"`
	// Optional. A list of tags that are added to the canary revision while the canary phase is in progress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#canary_revision_tags GoogleClouddeployDeliveryPipeline#canary_revision_tags}
	CanaryRevisionTags *[]*string `field:"optional" json:"canaryRevisionTags" yaml:"canaryRevisionTags"`
	// Optional. A list of tags that are added to the prior revision while the canary phase is in progress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#prior_revision_tags GoogleClouddeployDeliveryPipeline#prior_revision_tags}
	PriorRevisionTags *[]*string `field:"optional" json:"priorRevisionTags" yaml:"priorRevisionTags"`
	// Optional. A list of tags that are added to the final stable revision when the stable phase is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#stable_revision_tags GoogleClouddeployDeliveryPipeline#stable_revision_tags}
	StableRevisionTags *[]*string `field:"optional" json:"stableRevisionTags" yaml:"stableRevisionTags"`
}

