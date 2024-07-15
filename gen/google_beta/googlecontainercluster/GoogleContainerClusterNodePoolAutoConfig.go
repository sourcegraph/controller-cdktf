package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfig struct {
	// network_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#network_tags GoogleContainerCluster#network_tags}
	NetworkTags *GoogleContainerClusterNodePoolAutoConfigNetworkTags `field:"optional" json:"networkTags" yaml:"networkTags"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#resource_manager_tags GoogleContainerCluster#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

