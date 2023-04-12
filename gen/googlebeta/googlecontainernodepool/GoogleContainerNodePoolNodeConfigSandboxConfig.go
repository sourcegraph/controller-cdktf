package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#sandbox_type GoogleContainerNodePool#sandbox_type}
	SandboxType *string `field:"required" json:"sandboxType" yaml:"sandboxType"`
}

