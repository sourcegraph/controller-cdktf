package registrymodule


type RegistryModuleVcsRepo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#display_identifier RegistryModule#display_identifier}.
	DisplayIdentifier *string `field:"required" json:"displayIdentifier" yaml:"displayIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#identifier RegistryModule#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#github_app_installation_id RegistryModule#github_app_installation_id}.
	GithubAppInstallationId *string `field:"optional" json:"githubAppInstallationId" yaml:"githubAppInstallationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/registry_module#oauth_token_id RegistryModule#oauth_token_id}.
	OauthTokenId *string `field:"optional" json:"oauthTokenId" yaml:"oauthTokenId"`
}

