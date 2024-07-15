package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions struct {
	// Type of the redirect action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#type GoogleComputeSecurityPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#target GoogleComputeSecurityPolicy#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

