package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy struct {
	// The BackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_url_map#backend_service GoogleComputeUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

