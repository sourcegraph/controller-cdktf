package googlecomputebackendservice


type GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicy struct {
	// zonal_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_backend_service#zonal_affinity GoogleComputeBackendService#zonal_affinity}
	ZonalAffinity *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity `field:"optional" json:"zonalAffinity" yaml:"zonalAffinity"`
}

