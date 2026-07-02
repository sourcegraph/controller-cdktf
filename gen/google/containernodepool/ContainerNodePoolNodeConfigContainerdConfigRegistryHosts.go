package containernodepool

type ContainerNodePoolNodeConfigContainerdConfigRegistryHosts struct {
	// Defines the host name of the registry server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_node_pool#server ContainerNodePool#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_node_pool#hosts ContainerNodePool#hosts}
	Hosts any `field:"optional" json:"hosts" yaml:"hosts"`
}
