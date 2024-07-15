package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#sandbox_type GoogleContainerCluster#sandbox_type}
	SandboxType *string `field:"required" json:"sandboxType" yaml:"sandboxType"`
}

