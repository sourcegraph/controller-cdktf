package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfigRegistryHosts struct {
	// Defines the host name of the registry server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#server GoogleContainerNodePool#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#hosts GoogleContainerNodePool#hosts}
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
}

