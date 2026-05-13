package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigContainerdConfigRegistryHostsHostsHeader struct {
	// Configures the header key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#key GoogleContainerNodePool#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Configures the header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#value GoogleContainerNodePool#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

