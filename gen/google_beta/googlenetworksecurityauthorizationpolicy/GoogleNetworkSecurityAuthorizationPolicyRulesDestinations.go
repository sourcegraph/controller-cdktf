package googlenetworksecurityauthorizationpolicy


type GoogleNetworkSecurityAuthorizationPolicyRulesDestinations struct {
	// List of host names to match.
	//
	// Matched against the ":authority" header in http requests. At least one host should match. Each host can be an exact match, or a prefix match (example "mydomain.*") or a suffix match (example "*.myorg.com") or a presence (any) match "*".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#hosts GoogleNetworkSecurityAuthorizationPolicy#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// A list of HTTP methods to match.
	//
	// At least one method should match. Should not be set for gRPC services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#methods GoogleNetworkSecurityAuthorizationPolicy#methods}
	Methods *[]*string `field:"required" json:"methods" yaml:"methods"`
	// List of destination ports to match. At least one port should match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#ports GoogleNetworkSecurityAuthorizationPolicy#ports}
	Ports *[]*float64 `field:"required" json:"ports" yaml:"ports"`
	// http_header_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#http_header_match GoogleNetworkSecurityAuthorizationPolicy#http_header_match}
	HttpHeaderMatch *GoogleNetworkSecurityAuthorizationPolicyRulesDestinationsHttpHeaderMatch `field:"optional" json:"httpHeaderMatch" yaml:"httpHeaderMatch"`
}

