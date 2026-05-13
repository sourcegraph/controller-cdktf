package googlecontainercluster


type GoogleContainerClusterNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor'). Deprecated in favor of type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#sandbox_type GoogleContainerCluster#sandbox_type}
	SandboxType *string `field:"optional" json:"sandboxType" yaml:"sandboxType"`
	// Type of the sandbox to use for the node (e.g. 'GVISOR').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#type GoogleContainerCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

