package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy struct {
	// The full or partial URL to the RegionBackendService resource being mirrored to.
	//
	// The backend service configured for a mirroring policy must reference backends that are of the same type as the original backend service matched in the URL map.
	// Serverless NEG backends are not currently supported as a mirrored backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#backend_service GoogleComputeRegionUrlMap#backend_service}
	BackendService *string `field:"optional" json:"backendService" yaml:"backendService"`
	// The percentage of requests to be mirrored to backendService. The value must be between 0.0 and 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_url_map#mirror_percent GoogleComputeRegionUrlMap#mirror_percent}
	MirrorPercent *float64 `field:"optional" json:"mirrorPercent" yaml:"mirrorPercent"`
}

