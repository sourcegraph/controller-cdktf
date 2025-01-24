package containercluster


type ContainerClusterNodePoolQueuedProvisioning struct {
	// Whether nodes in this node pool are obtainable solely through the ProvisioningRequest API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.45.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

