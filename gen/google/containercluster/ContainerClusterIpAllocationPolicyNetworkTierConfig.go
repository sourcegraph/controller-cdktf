package containercluster


type ContainerClusterIpAllocationPolicyNetworkTierConfig struct {
	// Network tier configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#network_tier ContainerCluster#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
}

