package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRule struct {
	// Action to take when match matches the request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#action GoogleComputeSecurityPolicy#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#match GoogleComputeSecurityPolicy#match}
	Match *GoogleComputeSecurityPolicyRuleMatch `field:"required" json:"match" yaml:"match"`
	// An unique positive integer indicating the priority of evaluation for a rule.
	//
	// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#priority GoogleComputeSecurityPolicy#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// An optional description of this rule. Max size is 64.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#description GoogleComputeSecurityPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When set to true, the action specified above is not enforced.
	//
	// Stackdriver logs for requests that trigger a preview action are annotated as such.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#preview GoogleComputeSecurityPolicy#preview}
	Preview interface{} `field:"optional" json:"preview" yaml:"preview"`
	// rate_limit_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#rate_limit_options GoogleComputeSecurityPolicy#rate_limit_options}
	RateLimitOptions *GoogleComputeSecurityPolicyRuleRateLimitOptions `field:"optional" json:"rateLimitOptions" yaml:"rateLimitOptions"`
	// redirect_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#redirect_options GoogleComputeSecurityPolicy#redirect_options}
	RedirectOptions *GoogleComputeSecurityPolicyRuleRedirectOptions `field:"optional" json:"redirectOptions" yaml:"redirectOptions"`
}

