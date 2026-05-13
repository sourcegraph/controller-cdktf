package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig struct {
	// boot_disk_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#boot_disk_profile GoogleContainerNodePool#boot_disk_profile}
	BootDiskProfile *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile `field:"optional" json:"bootDiskProfile" yaml:"bootDiskProfile"`
	// dedicated_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#dedicated_local_ssd_profile GoogleContainerNodePool#dedicated_local_ssd_profile}
	DedicatedLocalSsdProfile *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile `field:"optional" json:"dedicatedLocalSsdProfile" yaml:"dedicatedLocalSsdProfile"`
	// Enables or disables swap for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#enabled GoogleContainerNodePool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#encryption_config GoogleContainerNodePool#encryption_config}
	EncryptionConfig *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// ephemeral_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#ephemeral_local_ssd_profile GoogleContainerNodePool#ephemeral_local_ssd_profile}
	EphemeralLocalSsdProfile *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile `field:"optional" json:"ephemeralLocalSsdProfile" yaml:"ephemeralLocalSsdProfile"`
}

