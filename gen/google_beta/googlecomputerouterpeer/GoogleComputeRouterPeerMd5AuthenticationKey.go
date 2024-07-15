package googlecomputerouterpeer


type GoogleComputeRouterPeerMd5AuthenticationKey struct {
	// Value of the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#key GoogleComputeRouterPeer#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// [REQUIRED] Name used to identify the key.
	//
	// Must be unique within a router. Must be referenced by exactly one bgpPeer. Must comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router_peer#name GoogleComputeRouterPeer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

