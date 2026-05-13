package googlecomputeinterconnectgroup


type GoogleComputeInterconnectGroupIntent struct {
	// The reliability the user intends this group to be capable of, in terms of the Interconnect product SLAs.
	//
	// Possible values: ["PRODUCTION_NON_CRITICAL", "PRODUCTION_CRITICAL", "NO_SLA", "AVAILABILITY_SLA_UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_group#topology_capability GoogleComputeInterconnectGroup#topology_capability}
	TopologyCapability *string `field:"optional" json:"topologyCapability" yaml:"topologyCapability"`
}

