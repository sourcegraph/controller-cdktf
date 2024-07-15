package containercluster


type ContainerClusterNodeConfigConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

