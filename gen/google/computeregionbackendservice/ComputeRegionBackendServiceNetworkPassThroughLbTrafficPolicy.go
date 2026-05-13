package computeregionbackendservice


type ComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy struct {
	// zonal_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_backend_service#zonal_affinity ComputeRegionBackendService#zonal_affinity}
	ZonalAffinity *ComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity `field:"optional" json:"zonalAffinity" yaml:"zonalAffinity"`
}

