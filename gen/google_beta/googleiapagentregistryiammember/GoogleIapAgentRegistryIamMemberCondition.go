package googleiapagentregistryiammember


type GoogleIapAgentRegistryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_member#expression GoogleIapAgentRegistryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_member#title GoogleIapAgentRegistryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_agent_registry_iam_member#description GoogleIapAgentRegistryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

