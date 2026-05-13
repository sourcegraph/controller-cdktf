package googleiamworkloadidentitypooliammember


type GoogleIamWorkloadIdentityPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_member#expression GoogleIamWorkloadIdentityPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_member#title GoogleIamWorkloadIdentityPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_member#description GoogleIamWorkloadIdentityPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

