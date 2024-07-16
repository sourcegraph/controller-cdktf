package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfigNetworkTags struct {
	// List of network tags applied to auto-provisioned node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#tags GoogleContainerCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

