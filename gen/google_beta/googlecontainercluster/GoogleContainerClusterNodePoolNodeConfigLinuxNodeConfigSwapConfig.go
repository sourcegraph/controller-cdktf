package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig struct {
	// boot_disk_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#boot_disk_profile GoogleContainerCluster#boot_disk_profile}
	BootDiskProfile *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile `field:"optional" json:"bootDiskProfile" yaml:"bootDiskProfile"`
	// dedicated_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#dedicated_local_ssd_profile GoogleContainerCluster#dedicated_local_ssd_profile}
	DedicatedLocalSsdProfile *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile `field:"optional" json:"dedicatedLocalSsdProfile" yaml:"dedicatedLocalSsdProfile"`
	// Enables or disables swap for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#encryption_config GoogleContainerCluster#encryption_config}
	EncryptionConfig *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// ephemeral_local_ssd_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#ephemeral_local_ssd_profile GoogleContainerCluster#ephemeral_local_ssd_profile}
	EphemeralLocalSsdProfile *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile `field:"optional" json:"ephemeralLocalSsdProfile" yaml:"ephemeralLocalSsdProfile"`
}

