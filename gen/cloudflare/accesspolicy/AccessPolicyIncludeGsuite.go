package accesspolicy


type AccessPolicyIncludeGsuite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#email AccessPolicy#email}.
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#identity_provider_id AccessPolicy#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}

