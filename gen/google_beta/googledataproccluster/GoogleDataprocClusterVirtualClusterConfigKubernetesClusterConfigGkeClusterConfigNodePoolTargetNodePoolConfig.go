package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig struct {
	// The list of Compute Engine zones where node pool nodes associated with a Dataproc on GKE virtual cluster will be located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#locations GoogleDataprocCluster#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#autoscaling GoogleDataprocCluster#autoscaling}
	Autoscaling *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#config GoogleDataprocCluster#config}
	Config *GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig `field:"optional" json:"config" yaml:"config"`
}

