package containercluster


type ContainerClusterNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'GVISOR').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#type ContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

