package googlecomputerouter


type GoogleComputeRouterBgpAdvertisedIpRanges struct {
	// The IP range to advertise. The value must be a CIDR-formatted string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router#range GoogleComputeRouter#range}
	Range *string `field:"required" json:"range" yaml:"range"`
	// User-specified description for the IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_router#description GoogleComputeRouter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

