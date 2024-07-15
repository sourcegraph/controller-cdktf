package googleclouddeploycustomtargettypeiambinding


type GoogleClouddeployCustomTargetTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_binding#expression GoogleClouddeployCustomTargetTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_binding#title GoogleClouddeployCustomTargetTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_binding#description GoogleClouddeployCustomTargetTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

