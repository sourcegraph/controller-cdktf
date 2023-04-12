package googlecontainercluster


type GoogleContainerClusterNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#sandbox_type GoogleContainerCluster#sandbox_type}
	SandboxType *string `field:"required" json:"sandboxType" yaml:"sandboxType"`
}

