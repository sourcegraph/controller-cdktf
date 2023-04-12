package googledataproccluster


type GoogleDataprocClusterClusterConfigAutoscalingConfig struct {
	// The autoscaling policy used by the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_cluster#policy_uri GoogleDataprocCluster#policy_uri}
	PolicyUri *string `field:"required" json:"policyUri" yaml:"policyUri"`
}

