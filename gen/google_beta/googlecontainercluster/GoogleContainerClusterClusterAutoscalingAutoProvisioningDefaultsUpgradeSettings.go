package googlecontainercluster


type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings struct {
	// blue_green_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#blue_green_settings GoogleContainerCluster#blue_green_settings}
	BlueGreenSettings *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings `field:"optional" json:"blueGreenSettings" yaml:"blueGreenSettings"`
	// The maximum number of nodes that can be created beyond the current size of the node pool during the upgrade process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#max_surge GoogleContainerCluster#max_surge}
	MaxSurge *float64 `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// The maximum number of nodes that can be simultaneously unavailable during the upgrade process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#max_unavailable GoogleContainerCluster#max_unavailable}
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Update strategy of the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#strategy GoogleContainerCluster#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

