package googleapigeeenvironmentiammember


type GoogleApigeeEnvironmentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_member#expression GoogleApigeeEnvironmentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_member#title GoogleApigeeEnvironmentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_apigee_environment_iam_member#description GoogleApigeeEnvironmentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

