package googlesiteverificationwebresource


type GoogleSiteVerificationWebResourceSite struct {
	// The site identifier.
	//
	// If the type is set to SITE, the identifier is a URL. If the type is
	// set to INET_DOMAIN, the identifier is a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_site_verification_web_resource#identifier GoogleSiteVerificationWebResource#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The type of resource to be verified. Possible values: ["INET_DOMAIN", "SITE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_site_verification_web_resource#type GoogleSiteVerificationWebResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

