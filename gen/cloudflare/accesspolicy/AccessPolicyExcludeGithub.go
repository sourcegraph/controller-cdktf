package accesspolicy


type AccessPolicyExcludeGithub struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#identity_provider_id AccessPolicy#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#name AccessPolicy#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#teams AccessPolicy#teams}.
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

