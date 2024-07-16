package googleclouddeploycustomtargettypeiammember


type GoogleClouddeployCustomTargetTypeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_member#expression GoogleClouddeployCustomTargetTypeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_member#title GoogleClouddeployCustomTargetTypeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_custom_target_type_iam_member#description GoogleClouddeployCustomTargetTypeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

