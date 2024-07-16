package googleclouddeploydeliverypipelineiambinding


type GoogleClouddeployDeliveryPipelineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_binding#expression GoogleClouddeployDeliveryPipelineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_binding#title GoogleClouddeployDeliveryPipelineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_delivery_pipeline_iam_binding#description GoogleClouddeployDeliveryPipelineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

