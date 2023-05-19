package accessgroup


type AccessGroupExcludeGsuite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#email AccessGroup#email}.
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_group#identity_provider_id AccessGroup#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}

