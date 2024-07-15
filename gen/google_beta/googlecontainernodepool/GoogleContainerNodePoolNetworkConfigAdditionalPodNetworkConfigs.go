package googlecontainernodepool


type GoogleContainerNodePoolNetworkConfigAdditionalPodNetworkConfigs struct {
	// The maximum number of pods per node which use this pod network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#max_pods_per_node GoogleContainerNodePool#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// The name of the secondary range on the subnet which provides IP address for this pod range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#secondary_pod_range GoogleContainerNodePool#secondary_pod_range}
	SecondaryPodRange *string `field:"optional" json:"secondaryPodRange" yaml:"secondaryPodRange"`
	// Name of the subnetwork where the additional pod network belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#subnetwork GoogleContainerNodePool#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

