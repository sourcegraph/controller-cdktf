package googlecomposerenvironment


type GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfig struct {
	// Whether or not master authorized networks is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#enabled GoogleComposerEnvironment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// cidr_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#cidr_blocks GoogleComposerEnvironment#cidr_blocks}
	CidrBlocks interface{} `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
}

