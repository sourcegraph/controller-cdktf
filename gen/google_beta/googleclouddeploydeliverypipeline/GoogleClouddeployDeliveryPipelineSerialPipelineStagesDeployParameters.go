package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters struct {
	// Required. Values are deploy parameters in key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#values GoogleClouddeployDeliveryPipeline#values}
	Values *map[string]*string `field:"required" json:"values" yaml:"values"`
	// Optional.
	//
	// Deploy parameters are applied to targets with match labels. If unspecified, deploy parameters are applied to all targets (including child targets of a multi-target).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline#match_target_labels GoogleClouddeployDeliveryPipeline#match_target_labels}
	MatchTargetLabels *map[string]*string `field:"optional" json:"matchTargetLabels" yaml:"matchTargetLabels"`
}

