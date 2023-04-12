package googlecomputeinstance


type GoogleComputeInstanceNetworkPerformanceConfig struct {
	// The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance#total_egress_bandwidth_tier GoogleComputeInstance#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

