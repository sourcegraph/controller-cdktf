package googlecomposerenvironment


type GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocks struct {
	// cidr_block must be specified in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#cidr_block GoogleComposerEnvironment#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// display_name is a field for users to identify CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#display_name GoogleComposerEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

