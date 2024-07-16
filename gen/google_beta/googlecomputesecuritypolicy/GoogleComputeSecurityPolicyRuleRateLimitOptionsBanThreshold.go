package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#count GoogleComputeSecurityPolicy#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#interval_sec GoogleComputeSecurityPolicy#interval_sec}
	IntervalSec *float64 `field:"required" json:"intervalSec" yaml:"intervalSec"`
}

