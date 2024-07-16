package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleHeaderActionRequestHeadersToAdds struct {
	// The name of the header to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#header_name GoogleComputeSecurityPolicy#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value to set the named header to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#header_value GoogleComputeSecurityPolicy#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

