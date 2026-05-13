package googlecontainernodepool


type GoogleContainerNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy struct {
	// Time in seconds to wait after cordoning the blue pool before draining the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#wait_for_drain_duration GoogleContainerNodePool#wait_for_drain_duration}
	WaitForDrainDuration *string `field:"optional" json:"waitForDrainDuration" yaml:"waitForDrainDuration"`
}

