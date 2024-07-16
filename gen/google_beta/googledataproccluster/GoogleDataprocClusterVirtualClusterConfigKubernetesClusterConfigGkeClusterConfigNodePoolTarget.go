package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget struct {
	// The target GKE node pool. Format: 'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{nodePool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_pool GoogleDataprocCluster#node_pool}
	NodePool *string `field:"required" json:"nodePool" yaml:"nodePool"`
	// The roles associated with the GKE node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#roles GoogleDataprocCluster#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_pool_config GoogleDataprocCluster#node_pool_config}
	NodePoolConfig *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig `field:"optional" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

