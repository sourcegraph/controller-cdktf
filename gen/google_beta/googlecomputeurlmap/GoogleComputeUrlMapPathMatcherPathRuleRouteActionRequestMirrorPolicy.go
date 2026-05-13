package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy struct {
	// The BackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#backend_service GoogleComputeUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
	// The percentage of requests to be mirrored to backendService. The value must be between 0.0 and 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_url_map#mirror_percent GoogleComputeUrlMap#mirror_percent}
	MirrorPercent *float64 `field:"optional" json:"mirrorPercent" yaml:"mirrorPercent"`
}

