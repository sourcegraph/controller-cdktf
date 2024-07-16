package identityplatformconfig


type IdentityPlatformConfigClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/identity_platform_config#permissions IdentityPlatformConfig#permissions}
	Permissions *IdentityPlatformConfigClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

