package policyset


type PolicySetVcsRepo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#identifier PolicySet#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#branch PolicySet#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#github_app_installation_id PolicySet#github_app_installation_id}.
	GithubAppInstallationId *string `field:"optional" json:"githubAppInstallationId" yaml:"githubAppInstallationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#ingress_submodules PolicySet#ingress_submodules}.
	IngressSubmodules interface{} `field:"optional" json:"ingressSubmodules" yaml:"ingressSubmodules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#oauth_token_id PolicySet#oauth_token_id}.
	OauthTokenId *string `field:"optional" json:"oauthTokenId" yaml:"oauthTokenId"`
}

