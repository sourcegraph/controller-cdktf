package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleMatchExprOptionsA struct {
	// recaptcha_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#recaptcha_options GoogleComputeSecurityPolicyRuleA#recaptcha_options}
	RecaptchaOptions *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA `field:"required" json:"recaptchaOptions" yaml:"recaptchaOptions"`
}

