package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfigNetworkTags struct {
	// List of network tags applied to auto-provisioned node pools.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#tags GoogleContainerCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

