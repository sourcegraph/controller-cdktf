package containercluster


type ContainerClusterNodePoolNodeConfigContainerdConfigRegistryHostsHostsClient struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#cert ContainerCluster#cert}
	Cert *ContainerClusterNodePoolNodeConfigContainerdConfigRegistryHostsHostsClientCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#key ContainerCluster#key}
	Key *ContainerClusterNodePoolNodeConfigContainerdConfigRegistryHostsHostsClientKey `field:"optional" json:"key" yaml:"key"`
}

