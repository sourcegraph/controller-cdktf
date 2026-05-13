package containercluster


type ContainerClusterNodeConfigKubeletConfigMemoryManager struct {
	// The Memory Manager policy to use.
	//
	// This policy guides how memory and hugepages are allocated and managed for pods on the node, influencing NUMA affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#policy ContainerCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

