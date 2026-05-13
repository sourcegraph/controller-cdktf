package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsClient struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#cert ContainerNodePool#cert}
	Cert *ContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsClientCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#key ContainerNodePool#key}
	Key *ContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsClientKey `field:"optional" json:"key" yaml:"key"`
}

