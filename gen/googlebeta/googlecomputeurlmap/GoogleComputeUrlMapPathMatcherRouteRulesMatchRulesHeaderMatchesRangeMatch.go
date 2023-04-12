package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch struct {
	// The end of the range (exclusive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#range_end GoogleComputeUrlMap#range_end}
	RangeEnd *float64 `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// The start of the range (inclusive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#range_start GoogleComputeUrlMap#range_start}
	RangeStart *float64 `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

