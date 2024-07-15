package googlecontainernodepool


type GoogleContainerNodePoolNetworkConfigNetworkPerformanceConfig struct {
	// Specifies the total network bandwidth tier for the NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#total_egress_bandwidth_tier GoogleContainerNodePool#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

