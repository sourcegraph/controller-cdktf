package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicySelectorsDeliveryPipeline struct {
	// Optional.
	//
	// ID of the DeliveryPipeline. The value of this field could be one of the following:
	// - The last segment of a pipeline name
	// - "*", all delivery pipelines in a location
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#id GoogleClouddeployDeployPolicy#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// DeliveryPipeline labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddeploy_deploy_policy#labels GoogleClouddeployDeployPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

