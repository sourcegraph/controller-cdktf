package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigWorkloadMetadataConfig struct {
	// Mode is the configuration for how to expose metadata to workloads running on the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#mode GoogleContainerNodePool#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

