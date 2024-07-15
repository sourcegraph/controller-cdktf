package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#abort GoogleComputeUrlMap#abort}
	Abort *GoogleComputeUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_url_map#delay GoogleComputeUrlMap#delay}
	Delay *GoogleComputeUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

