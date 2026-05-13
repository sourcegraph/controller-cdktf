package googlecomputeorganizationsecuritypolicyrule


type GoogleComputeOrganizationSecurityPolicyRuleHeaderActionRequestHeadersToAdds struct {
	// The name of the header to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#header_name GoogleComputeOrganizationSecurityPolicyRule#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value to set the named header to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_organization_security_policy_rule#header_value GoogleComputeOrganizationSecurityPolicyRule#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

