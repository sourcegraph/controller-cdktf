package containercluster


type ContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClient struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#cert ContainerCluster#cert}
	Cert *ContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClientCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#key ContainerCluster#key}
	Key *ContainerClusterNodeConfigContainerdConfigRegistryHostsHostsClientKey `field:"optional" json:"key" yaml:"key"`
}

