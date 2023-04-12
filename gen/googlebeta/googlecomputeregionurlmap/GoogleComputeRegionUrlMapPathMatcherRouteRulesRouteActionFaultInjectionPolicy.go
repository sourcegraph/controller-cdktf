package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_url_map#abort GoogleComputeRegionUrlMap#abort}
	Abort *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_url_map#delay GoogleComputeRegionUrlMap#delay}
	Delay *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

