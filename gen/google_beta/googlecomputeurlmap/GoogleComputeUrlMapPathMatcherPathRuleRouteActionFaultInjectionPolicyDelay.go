package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay struct {
	// fixed_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#fixed_delay GoogleComputeUrlMap#fixed_delay}
	FixedDelay *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay `field:"required" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.
	//
	// The value must be between 0.0 and
	// 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#percentage GoogleComputeUrlMap#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

