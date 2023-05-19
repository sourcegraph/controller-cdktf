package googledataproccluster


type GoogleDataprocClusterClusterConfigAutoscalingConfig struct {
	// The autoscaling policy used by the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataproc_cluster#policy_uri GoogleDataprocCluster#policy_uri}
	PolicyUri *string `field:"required" json:"policyUri" yaml:"policyUri"`
}

