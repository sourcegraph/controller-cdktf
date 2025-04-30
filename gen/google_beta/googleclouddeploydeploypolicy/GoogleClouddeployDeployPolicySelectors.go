package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicySelectors struct {
	// delivery_pipeline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#delivery_pipeline GoogleClouddeployDeployPolicy#delivery_pipeline}
	DeliveryPipeline *GoogleClouddeployDeployPolicySelectorsDeliveryPipeline `field:"optional" json:"deliveryPipeline" yaml:"deliveryPipeline"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_clouddeploy_deploy_policy#target GoogleClouddeployDeployPolicy#target}
	Target *GoogleClouddeployDeployPolicySelectorsTarget `field:"optional" json:"target" yaml:"target"`
}

