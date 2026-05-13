package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceSubsetting struct {
	// The algorithm used for subsetting. Possible values: ["CONSISTENT_HASH_SUBSETTING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#policy GoogleComputeRegionBackendService#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The number of backends per backend group assigned to each proxy instance or each service mesh client.
	//
	// An input parameter to the CONSISTENT_HASH_SUBSETTING algorithm. Can only be set if policy is set to
	// CONSISTENT_HASH_SUBSETTING. Can only be set if load balancing scheme is INTERNAL_MANAGED or INTERNAL_SELF_MANAGED.
	// subsetSize is optional for Internal HTTP(S) load balancing and required for Traffic Director.
	// If you do not provide this value, Cloud Load Balancing will calculate it dynamically to optimize the number
	// of proxies/clients visible to each backend and vice versa.
	// Must be greater than 0. If subsetSize is larger than the number of backends/endpoints, then subsetting is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#subset_size GoogleComputeRegionBackendService#subset_size}
	SubsetSize *float64 `field:"optional" json:"subsetSize" yaml:"subsetSize"`
}

