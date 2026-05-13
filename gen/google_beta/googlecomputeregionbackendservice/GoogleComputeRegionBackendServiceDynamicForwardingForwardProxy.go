package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceDynamicForwardingForwardProxy struct {
	// A boolean flag enabling dynamic forwarding proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#enabled GoogleComputeRegionBackendService#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Determines the dynamic forwarding proxy mode Possible values: ["DIRECT_FORWARDING", "CLOUD_RUN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#proxy_mode GoogleComputeRegionBackendService#proxy_mode}
	ProxyMode *string `field:"required" json:"proxyMode" yaml:"proxyMode"`
}

