package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsHeader struct {
	// Configures the header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#key ContainerNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Configures the header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#value ContainerNodePool#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

