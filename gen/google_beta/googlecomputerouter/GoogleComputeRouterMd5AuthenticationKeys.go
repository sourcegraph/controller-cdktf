package googlecomputerouter


type GoogleComputeRouterMd5AuthenticationKeys struct {
	// Value of the key used for MD5 authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router#key GoogleComputeRouter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name used to identify the key.
	//
	// Must be unique within a router.
	// Must be referenced by exactly one bgpPeer. Must comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router#name GoogleComputeRouter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

