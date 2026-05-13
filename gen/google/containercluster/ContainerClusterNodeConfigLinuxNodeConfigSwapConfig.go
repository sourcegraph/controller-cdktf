package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigSwapConfig struct {
	// boot_disk_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#boot_disk_profile ContainerCluster#boot_disk_profile}
	BootDiskProfile *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile `field:"optional" json:"bootDiskProfile" yaml:"bootDiskProfile"`
	// dedicated_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#dedicated_local_ssd_profile ContainerCluster#dedicated_local_ssd_profile}
	DedicatedLocalSsdProfile *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile `field:"optional" json:"dedicatedLocalSsdProfile" yaml:"dedicatedLocalSsdProfile"`
	// Enables or disables swap for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#encryption_config ContainerCluster#encryption_config}
	EncryptionConfig *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// ephemeral_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#ephemeral_local_ssd_profile ContainerCluster#ephemeral_local_ssd_profile}
	EphemeralLocalSsdProfile *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile `field:"optional" json:"ephemeralLocalSsdProfile" yaml:"ephemeralLocalSsdProfile"`
}

