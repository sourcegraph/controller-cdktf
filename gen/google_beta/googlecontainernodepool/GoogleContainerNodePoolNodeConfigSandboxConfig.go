package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#sandbox_type GoogleContainerNodePool#sandbox_type}
	SandboxType *string `field:"required" json:"sandboxType" yaml:"sandboxType"`
}

