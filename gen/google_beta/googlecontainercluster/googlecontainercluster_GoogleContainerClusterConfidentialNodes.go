package googlecontainercluster


type GoogleContainerClusterConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

