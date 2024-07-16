package googleedgecontainercluster


type GoogleEdgecontainerClusterControlPlaneRemote struct {
	// Name of the Google Distributed Cloud Edge zones where this node pool will be created. For example: 'us-central1-edge-customer-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#node_location GoogleEdgecontainerCluster#node_location}
	NodeLocation *string `field:"optional" json:"nodeLocation" yaml:"nodeLocation"`
}

