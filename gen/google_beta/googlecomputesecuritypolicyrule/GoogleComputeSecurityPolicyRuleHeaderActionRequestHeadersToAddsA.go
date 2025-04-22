package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleHeaderActionRequestHeadersToAddsA struct {
	// The name of the header to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_security_policy_rule#header_name GoogleComputeSecurityPolicyRuleA#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value to set the named header to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_security_policy_rule#header_value GoogleComputeSecurityPolicyRuleA#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

