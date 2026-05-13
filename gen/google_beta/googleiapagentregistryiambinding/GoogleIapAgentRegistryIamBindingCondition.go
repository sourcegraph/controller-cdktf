package googleiapagentregistryiambinding


type GoogleIapAgentRegistryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_binding#expression GoogleIapAgentRegistryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_binding#title GoogleIapAgentRegistryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_binding#description GoogleIapAgentRegistryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

