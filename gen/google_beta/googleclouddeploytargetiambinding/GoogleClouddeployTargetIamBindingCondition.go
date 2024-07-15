package googleclouddeploytargetiambinding


type GoogleClouddeployTargetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_binding#expression GoogleClouddeployTargetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_binding#title GoogleClouddeployTargetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_binding#description GoogleClouddeployTargetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

