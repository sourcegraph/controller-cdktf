package workspace


type WorkspaceVcsRepo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/workspace#identifier Workspace#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/workspace#oauth_token_id Workspace#oauth_token_id}.
	OauthTokenId *string `field:"required" json:"oauthTokenId" yaml:"oauthTokenId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/workspace#branch Workspace#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/workspace#ingress_submodules Workspace#ingress_submodules}.
	IngressSubmodules interface{} `field:"optional" json:"ingressSubmodules" yaml:"ingressSubmodules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/workspace#tags_regex Workspace#tags_regex}.
	TagsRegex *string `field:"optional" json:"tagsRegex" yaml:"tagsRegex"`
}

