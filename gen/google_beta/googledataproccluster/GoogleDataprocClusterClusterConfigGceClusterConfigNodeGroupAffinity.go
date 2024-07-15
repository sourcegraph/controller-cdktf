package googledataproccluster


type GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity struct {
	// The URI of a sole-tenant that the cluster will be created on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_group_uri GoogleDataprocCluster#node_group_uri}
	NodeGroupUri *string `field:"required" json:"nodeGroupUri" yaml:"nodeGroupUri"`
}

