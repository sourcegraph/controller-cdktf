package googlecontainercluster


type GoogleContainerClusterIpAllocationPolicyNetworkTierConfig struct {
	// Network tier configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#network_tier GoogleContainerCluster#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
}

