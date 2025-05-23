package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy struct {
	// The RegionBackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_region_url_map#backend_service GoogleComputeRegionUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

