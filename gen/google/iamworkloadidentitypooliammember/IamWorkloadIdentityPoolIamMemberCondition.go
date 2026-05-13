package iamworkloadidentitypooliammember


type IamWorkloadIdentityPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_iam_member#expression IamWorkloadIdentityPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_iam_member#title IamWorkloadIdentityPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_iam_member#description IamWorkloadIdentityPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

