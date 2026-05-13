package containernodepool


type ContainerNodePoolNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#private_registry_access_config ContainerNodePool#private_registry_access_config}
	PrivateRegistryAccessConfig *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
	// registry_hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#registry_hosts ContainerNodePool#registry_hosts}
	RegistryHosts interface{} `field:"optional" json:"registryHosts" yaml:"registryHosts"`
	// writable_cgroups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#writable_cgroups ContainerNodePool#writable_cgroups}
	WritableCgroups *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups `field:"optional" json:"writableCgroups" yaml:"writableCgroups"`
}

