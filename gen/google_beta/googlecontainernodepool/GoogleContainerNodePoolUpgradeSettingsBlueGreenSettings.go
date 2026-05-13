package googlecontainernodepool


type GoogleContainerNodePoolUpgradeSettingsBlueGreenSettings struct {
	// autoscaled_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#autoscaled_rollout_policy GoogleContainerNodePool#autoscaled_rollout_policy}
	AutoscaledRolloutPolicy *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy `field:"optional" json:"autoscaledRolloutPolicy" yaml:"autoscaledRolloutPolicy"`
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#node_pool_soak_duration GoogleContainerNodePool#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#standard_rollout_policy GoogleContainerNodePool#standard_rollout_policy}
	StandardRolloutPolicy *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"optional" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
}

