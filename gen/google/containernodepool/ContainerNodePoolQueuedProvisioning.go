package containernodepool


type ContainerNodePoolQueuedProvisioning struct {
	// Whether nodes in this node pool are obtainable solely through the ProvisioningRequest API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

