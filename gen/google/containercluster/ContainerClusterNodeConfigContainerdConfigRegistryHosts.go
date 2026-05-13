package containercluster


type ContainerClusterNodeConfigContainerdConfigRegistryHosts struct {
	// Defines the host name of the registry server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#server ContainerCluster#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#hosts ContainerCluster#hosts}
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
}

