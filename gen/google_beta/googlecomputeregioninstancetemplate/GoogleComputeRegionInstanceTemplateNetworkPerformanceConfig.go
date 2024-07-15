package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateNetworkPerformanceConfig struct {
	// The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_instance_template#total_egress_bandwidth_tier GoogleComputeRegionInstanceTemplate#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

