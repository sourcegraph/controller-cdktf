package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#count GoogleComputeSecurityPolicy#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_security_policy#interval_sec GoogleComputeSecurityPolicy#interval_sec}
	IntervalSec *float64 `field:"required" json:"intervalSec" yaml:"intervalSec"`
}

