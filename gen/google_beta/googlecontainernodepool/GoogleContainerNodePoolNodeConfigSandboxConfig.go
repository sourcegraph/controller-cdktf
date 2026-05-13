package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor'). Deprecated in favor of type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#sandbox_type GoogleContainerNodePool#sandbox_type}
	SandboxType *string `field:"optional" json:"sandboxType" yaml:"sandboxType"`
	// Type of the sandbox to use for the node (e.g. 'GVISOR').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#type GoogleContainerNodePool#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

