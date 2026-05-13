package googlecontainercluster


type GoogleContainerClusterIpAllocationPolicyAdditionalIpRangesConfig struct {
	// Name of the subnetwork. This can be the full path of the subnetwork or just the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#subnetwork GoogleContainerCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
	// List of secondary ranges names within this subnetwork that can be used for pod IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#pod_ipv4_range_names GoogleContainerCluster#pod_ipv4_range_names}
	PodIpv4RangeNames *[]*string `field:"optional" json:"podIpv4RangeNames" yaml:"podIpv4RangeNames"`
	// Status of the subnetwork, If in draining status, subnet will not be selected for new node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#status GoogleContainerCluster#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

