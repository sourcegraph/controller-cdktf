package googlecomputesecuritypolicyrule


type GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#count GoogleComputeSecurityPolicyRuleA#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy_rule#interval_sec GoogleComputeSecurityPolicyRuleA#interval_sec}
	IntervalSec *float64 `field:"optional" json:"intervalSec" yaml:"intervalSec"`
}

