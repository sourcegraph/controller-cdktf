package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_region_url_map#abort GoogleComputeRegionUrlMap#abort}
	Abort *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_region_url_map#delay GoogleComputeRegionUrlMap#delay}
	Delay *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

