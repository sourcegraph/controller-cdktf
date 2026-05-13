package googlecontainercluster


type GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettings struct {
	// autoscaled_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#autoscaled_rollout_policy GoogleContainerCluster#autoscaled_rollout_policy}
	AutoscaledRolloutPolicy *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy `field:"optional" json:"autoscaledRolloutPolicy" yaml:"autoscaledRolloutPolicy"`
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#node_pool_soak_duration GoogleContainerCluster#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#standard_rollout_policy GoogleContainerCluster#standard_rollout_policy}
	StandardRolloutPolicy *GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"optional" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
}

