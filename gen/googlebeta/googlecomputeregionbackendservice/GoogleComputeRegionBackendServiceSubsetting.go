package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceSubsetting struct {
	// The algorithm used for subsetting. Possible values: ["CONSISTENT_HASH_SUBSETTING"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_backend_service#policy GoogleComputeRegionBackendService#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

