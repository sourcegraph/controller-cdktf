package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#abort GoogleComputeRegionUrlMap#abort}
	Abort *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#delay GoogleComputeRegionUrlMap#delay}
	Delay *GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

