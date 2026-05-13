package containercluster


type ContainerClusterNodePoolNodeDrainConfig struct {
	// Whether to respect PodDisruptionBudget policy during node pool deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#respect_pdb_during_node_pool_deletion ContainerCluster#respect_pdb_during_node_pool_deletion}
	RespectPdbDuringNodePoolDeletion interface{} `field:"optional" json:"respectPdbDuringNodePoolDeletion" yaml:"respectPdbDuringNodePoolDeletion"`
}

