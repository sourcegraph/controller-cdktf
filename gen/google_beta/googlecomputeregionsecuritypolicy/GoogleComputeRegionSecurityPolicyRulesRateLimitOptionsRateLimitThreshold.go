package googlecomputeregionsecuritypolicy


type GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_security_policy#count GoogleComputeRegionSecurityPolicy#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_region_security_policy#interval_sec GoogleComputeRegionSecurityPolicy#interval_sec}
	IntervalSec *float64 `field:"optional" json:"intervalSec" yaml:"intervalSec"`
}

