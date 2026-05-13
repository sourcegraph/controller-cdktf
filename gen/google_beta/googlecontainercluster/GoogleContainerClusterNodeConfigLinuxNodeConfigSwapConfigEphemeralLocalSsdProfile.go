package googlecontainercluster


type GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile struct {
	// Specifies the size of the swap space in gibibytes (GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#swap_size_gib GoogleContainerCluster#swap_size_gib}
	SwapSizeGib *float64 `field:"optional" json:"swapSizeGib" yaml:"swapSizeGib"`
	// Specifies the size of the swap space as a percentage of the ephemeral local SSD capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#swap_size_percent GoogleContainerCluster#swap_size_percent}
	SwapSizePercent *float64 `field:"optional" json:"swapSizePercent" yaml:"swapSizePercent"`
}

