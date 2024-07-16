package googleaccesscontextmanageraccesspolicyiambinding


type GoogleAccessContextManagerAccessPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_policy_iam_binding#expression GoogleAccessContextManagerAccessPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_policy_iam_binding#title GoogleAccessContextManagerAccessPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_policy_iam_binding#description GoogleAccessContextManagerAccessPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

