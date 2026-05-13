package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigWritableCgroups struct {
	// Whether writable cgroups are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

