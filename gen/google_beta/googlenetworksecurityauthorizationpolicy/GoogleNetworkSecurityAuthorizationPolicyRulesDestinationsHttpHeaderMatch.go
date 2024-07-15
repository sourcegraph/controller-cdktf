package googlenetworksecurityauthorizationpolicy


type GoogleNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch struct {
	// The name of the HTTP header to match.
	//
	// For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#header_name GoogleNetworkSecurityAuthorizationPolicy#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header must match the regular expression specified in regexMatch.
	//
	// For regular expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript For matching against a port specified in the HTTP request, use a headerMatch with headerName set to Host and a regular expression that satisfies the RFC2616 Host header's port specifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#regex_match GoogleNetworkSecurityAuthorizationPolicy#regex_match}
	RegexMatch *string `field:"required" json:"regexMatch" yaml:"regexMatch"`
}

