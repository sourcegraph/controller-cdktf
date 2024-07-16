package googleclouddeploytargetiammember


type GoogleClouddeployTargetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_member#expression GoogleClouddeployTargetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_member#title GoogleClouddeployTargetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target_iam_member#description GoogleClouddeployTargetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

