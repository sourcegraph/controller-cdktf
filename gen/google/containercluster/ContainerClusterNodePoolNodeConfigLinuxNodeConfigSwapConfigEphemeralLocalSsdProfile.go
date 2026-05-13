package containercluster


type ContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile struct {
	// Specifies the size of the swap space in gibibytes (GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#swap_size_gib ContainerCluster#swap_size_gib}
	SwapSizeGib *float64 `field:"optional" json:"swapSizeGib" yaml:"swapSizeGib"`
	// Specifies the size of the swap space as a percentage of the ephemeral local SSD capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#swap_size_percent ContainerCluster#swap_size_percent}
	SwapSizePercent *float64 `field:"optional" json:"swapSizePercent" yaml:"swapSizePercent"`
}

