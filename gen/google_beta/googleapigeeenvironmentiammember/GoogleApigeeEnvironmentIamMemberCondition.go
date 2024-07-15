package googleapigeeenvironmentiammember


type GoogleApigeeEnvironmentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_environment_iam_member#expression GoogleApigeeEnvironmentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_environment_iam_member#title GoogleApigeeEnvironmentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_environment_iam_member#description GoogleApigeeEnvironmentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

