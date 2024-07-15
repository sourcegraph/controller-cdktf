package googlecontainernodepool


type GoogleContainerNodePoolUpgradeSettingsBlueGreenSettings struct {
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#standard_rollout_policy GoogleContainerNodePool#standard_rollout_policy}
	StandardRolloutPolicy *GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"required" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#node_pool_soak_duration GoogleContainerNodePool#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
}

