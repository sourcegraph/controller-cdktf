package accessorganization


type AccessOrganizationCustomPages struct {
	// The id of the forbidden page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_organization#forbidden AccessOrganization#forbidden}
	Forbidden *string `field:"optional" json:"forbidden" yaml:"forbidden"`
	// The id of the identity denied page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/access_organization#identity_denied AccessOrganization#identity_denied}
	IdentityDenied *string `field:"optional" json:"identityDenied" yaml:"identityDenied"`
}

