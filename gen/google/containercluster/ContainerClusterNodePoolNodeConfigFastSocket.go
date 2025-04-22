package containercluster


type ContainerClusterNodePoolNodeConfigFastSocket struct {
	// Whether or not NCCL Fast Socket is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

