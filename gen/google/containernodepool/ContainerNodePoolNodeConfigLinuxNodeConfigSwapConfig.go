package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig struct {
	// boot_disk_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#boot_disk_profile ContainerNodePool#boot_disk_profile}
	BootDiskProfile *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile `field:"optional" json:"bootDiskProfile" yaml:"bootDiskProfile"`
	// dedicated_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#dedicated_local_ssd_profile ContainerNodePool#dedicated_local_ssd_profile}
	DedicatedLocalSsdProfile *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile `field:"optional" json:"dedicatedLocalSsdProfile" yaml:"dedicatedLocalSsdProfile"`
	// Enables or disables swap for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#encryption_config ContainerNodePool#encryption_config}
	EncryptionConfig *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// ephemeral_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#ephemeral_local_ssd_profile ContainerNodePool#ephemeral_local_ssd_profile}
	EphemeralLocalSsdProfile *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile `field:"optional" json:"ephemeralLocalSsdProfile" yaml:"ephemeralLocalSsdProfile"`
}

