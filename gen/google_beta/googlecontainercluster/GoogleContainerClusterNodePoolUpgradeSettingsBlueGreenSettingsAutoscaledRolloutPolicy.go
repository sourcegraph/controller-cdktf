package googlecontainercluster


type GoogleContainerClusterNodePoolUpgradeSettingsBlueGreenSettingsAutoscaledRolloutPolicy struct {
	// Time in seconds to wait after cordoning the blue pool before draining the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#wait_for_drain_duration GoogleContainerCluster#wait_for_drain_duration}
	WaitForDrainDuration *string `field:"optional" json:"waitForDrainDuration" yaml:"waitForDrainDuration"`
}

