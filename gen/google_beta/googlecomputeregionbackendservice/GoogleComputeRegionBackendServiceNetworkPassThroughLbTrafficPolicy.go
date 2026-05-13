package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy struct {
	// zonal_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#zonal_affinity GoogleComputeRegionBackendService#zonal_affinity}
	ZonalAffinity *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity `field:"optional" json:"zonalAffinity" yaml:"zonalAffinity"`
}

