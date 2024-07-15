package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request. The value must be between 200 and 599 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#http_status GoogleComputeRegionUrlMap#http_status}
	HttpStatus *float64 `field:"required" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic (connections/operations/requests) which will be aborted as part of fault injection.
	//
	// The value must be between 0.0 and 100.0
	// inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#percentage GoogleComputeRegionUrlMap#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

