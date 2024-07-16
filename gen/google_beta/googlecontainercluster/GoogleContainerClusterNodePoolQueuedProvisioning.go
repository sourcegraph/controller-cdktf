package googlecontainercluster


type GoogleContainerClusterNodePoolQueuedProvisioning struct {
	// Whether nodes in this node pool are obtainable solely through the ProvisioningRequest API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

