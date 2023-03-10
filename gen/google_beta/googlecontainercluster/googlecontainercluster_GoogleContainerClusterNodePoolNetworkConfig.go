package googlecontainercluster


type GoogleContainerClusterNodePoolNetworkConfig struct {
	// The ID of the secondary range for pod IPs.
	//
	// If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#pod_range GoogleContainerCluster#pod_range}
	PodRange *string `field:"required" json:"podRange" yaml:"podRange"`
	// Whether to create a new range for pod IPs in this node pool.
	//
	// Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#create_pod_range GoogleContainerCluster#create_pod_range}
	CreatePodRange interface{} `field:"optional" json:"createPodRange" yaml:"createPodRange"`
	// The IP address range for pod IPs in this node pool.
	//
	// Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#pod_ipv4_cidr_block GoogleContainerCluster#pod_ipv4_cidr_block}
	PodIpv4CidrBlock *string `field:"optional" json:"podIpv4CidrBlock" yaml:"podIpv4CidrBlock"`
}

