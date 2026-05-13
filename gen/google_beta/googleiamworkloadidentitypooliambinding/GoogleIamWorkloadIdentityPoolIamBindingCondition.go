package googleiamworkloadidentitypooliambinding


type GoogleIamWorkloadIdentityPoolIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_binding#expression GoogleIamWorkloadIdentityPoolIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_binding#title GoogleIamWorkloadIdentityPoolIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workload_identity_pool_iam_binding#description GoogleIamWorkloadIdentityPoolIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

