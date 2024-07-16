package googlecontainercluster


type GoogleContainerClusterIpAllocationPolicyAdditionalPodRangesConfig struct {
	// Name for pod secondary ipv4 range which has the actual range defined ahead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#pod_range_names GoogleContainerCluster#pod_range_names}
	PodRangeNames *[]*string `field:"required" json:"podRangeNames" yaml:"podRangeNames"`
}

