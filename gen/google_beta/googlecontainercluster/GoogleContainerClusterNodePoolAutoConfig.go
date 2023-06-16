package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfig struct {
	// network_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#network_tags GoogleContainerCluster#network_tags}
	NetworkTags *GoogleContainerClusterNodePoolAutoConfigNetworkTags `field:"optional" json:"networkTags" yaml:"networkTags"`
}

