package computeregionbackendservice


type ComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity struct {
	// This field indicates whether zonal affinity is enabled or not. Default value: "ZONAL_AFFINITY_DISABLED" Possible values: ["ZONAL_AFFINITY_DISABLED", "ZONAL_AFFINITY_SPILL_CROSS_ZONE", "ZONAL_AFFINITY_STAY_WITHIN_ZONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_backend_service#spillover ComputeRegionBackendService#spillover}
	Spillover *string `field:"optional" json:"spillover" yaml:"spillover"`
	// The value of the field must be in [0, 1].
	//
	// When the ratio of the count of healthy backend endpoints in a zone
	// to the count of backend endpoints in that same zone is equal to or above this threshold, the load balancer
	// distributes new connections to all healthy endpoints in the local zone only. When the ratio of the count
	// of healthy backend endpoints in a zone to the count of backend endpoints in that same zone is below this
	// threshold, the load balancer distributes all new connections to all healthy endpoints across all zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_backend_service#spillover_ratio ComputeRegionBackendService#spillover_ratio}
	SpilloverRatio *float64 `field:"optional" json:"spilloverRatio" yaml:"spilloverRatio"`
}

