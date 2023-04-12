package googlebinaryauthorizationpolicy


type GoogleBinaryAuthorizationPolicyAdmissionWhitelistPatterns struct {
	// An image name pattern to whitelist, in the form 'registry/path/to/image'.
	//
	// This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_policy#name_pattern GoogleBinaryAuthorizationPolicy#name_pattern}
	NamePattern *string `field:"required" json:"namePattern" yaml:"namePattern"`
}

