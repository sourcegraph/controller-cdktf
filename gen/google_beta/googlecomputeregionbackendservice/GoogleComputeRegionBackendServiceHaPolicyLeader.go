package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceHaPolicyLeader struct {
	// A fully-qualified URL of the zonal Network Endpoint Group (NEG) that the leader is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#backend_group GoogleComputeRegionBackendService#backend_group}
	BackendGroup *string `field:"optional" json:"backendGroup" yaml:"backendGroup"`
	// network_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_region_backend_service#network_endpoint GoogleComputeRegionBackendService#network_endpoint}
	NetworkEndpoint *GoogleComputeRegionBackendServiceHaPolicyLeaderNetworkEndpoint `field:"optional" json:"networkEndpoint" yaml:"networkEndpoint"`
}

