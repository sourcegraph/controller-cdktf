package policyset


type PolicySetVcsRepo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/policy_set#identifier PolicySet#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/policy_set#oauth_token_id PolicySet#oauth_token_id}.
	OauthTokenId *string `field:"required" json:"oauthTokenId" yaml:"oauthTokenId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/policy_set#branch PolicySet#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/policy_set#ingress_submodules PolicySet#ingress_submodules}.
	IngressSubmodules interface{} `field:"optional" json:"ingressSubmodules" yaml:"ingressSubmodules"`
}

