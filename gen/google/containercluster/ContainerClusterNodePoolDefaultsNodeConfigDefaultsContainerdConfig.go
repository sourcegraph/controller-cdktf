package containercluster


type ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#private_registry_access_config ContainerCluster#private_registry_access_config}
	PrivateRegistryAccessConfig *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
	// registry_hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#registry_hosts ContainerCluster#registry_hosts}
	RegistryHosts interface{} `field:"optional" json:"registryHosts" yaml:"registryHosts"`
	// writable_cgroups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#writable_cgroups ContainerCluster#writable_cgroups}
	WritableCgroups *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigWritableCgroups `field:"optional" json:"writableCgroups" yaml:"writableCgroups"`
}

