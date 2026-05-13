package computeregionsecuritypolicyrule


type ComputeRegionSecurityPolicyRuleRateLimitOptionsRateLimitThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_security_policy_rule#count ComputeRegionSecurityPolicyRule#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_security_policy_rule#interval_sec ComputeRegionSecurityPolicyRule#interval_sec}
	IntervalSec *float64 `field:"optional" json:"intervalSec" yaml:"intervalSec"`
}

