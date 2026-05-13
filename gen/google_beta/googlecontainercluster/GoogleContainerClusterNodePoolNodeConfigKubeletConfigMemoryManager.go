package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigKubeletConfigMemoryManager struct {
	// The Memory Manager policy to use.
	//
	// This policy guides how memory and hugepages are allocated and managed for pods on the node, influencing NUMA affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#policy GoogleContainerCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

