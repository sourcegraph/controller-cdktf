package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigContainerdConfigRegistryHosts struct {
	// Defines the host name of the registry server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#server GoogleContainerCluster#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#hosts GoogleContainerCluster#hosts}
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
}

