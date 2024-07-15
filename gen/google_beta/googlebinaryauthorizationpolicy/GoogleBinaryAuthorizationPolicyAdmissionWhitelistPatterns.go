package googlebinaryauthorizationpolicy


type GoogleBinaryAuthorizationPolicyAdmissionWhitelistPatterns struct {
	// An image name pattern to whitelist, in the form 'registry/path/to/image'.
	//
	// This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#name_pattern GoogleBinaryAuthorizationPolicy#name_pattern}
	NamePattern *string `field:"required" json:"namePattern" yaml:"namePattern"`
}

