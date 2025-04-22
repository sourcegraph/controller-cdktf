package googlecomputerouterpeer


type GoogleComputeRouterPeerCustomLearnedIpRanges struct {
	// The IP range to learn. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router_peer#range GoogleComputeRouterPeer#range}
	Range *string `field:"required" json:"range" yaml:"range"`
}

