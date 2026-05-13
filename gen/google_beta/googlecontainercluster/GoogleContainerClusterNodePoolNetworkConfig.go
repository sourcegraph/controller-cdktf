package googlecontainercluster


type GoogleContainerClusterNodePoolNetworkConfig struct {
	// The accelerator network profile to use for this node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#accelerator_network_profile GoogleContainerCluster#accelerator_network_profile}
	AcceleratorNetworkProfile *string `field:"optional" json:"acceleratorNetworkProfile" yaml:"acceleratorNetworkProfile"`
	// additional_node_network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#additional_node_network_configs GoogleContainerCluster#additional_node_network_configs}
	AdditionalNodeNetworkConfigs interface{} `field:"optional" json:"additionalNodeNetworkConfigs" yaml:"additionalNodeNetworkConfigs"`
	// additional_pod_network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#additional_pod_network_configs GoogleContainerCluster#additional_pod_network_configs}
	AdditionalPodNetworkConfigs interface{} `field:"optional" json:"additionalPodNetworkConfigs" yaml:"additionalPodNetworkConfigs"`
	// Whether to create a new range for pod IPs in this node pool.
	//
	// Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#create_pod_range GoogleContainerCluster#create_pod_range}
	CreatePodRange interface{} `field:"optional" json:"createPodRange" yaml:"createPodRange"`
	// Whether nodes have internal IP addresses only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enable_private_nodes GoogleContainerCluster#enable_private_nodes}
	EnablePrivateNodes interface{} `field:"optional" json:"enablePrivateNodes" yaml:"enablePrivateNodes"`
	// network_performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#network_performance_config GoogleContainerCluster#network_performance_config}
	NetworkPerformanceConfig *GoogleContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig `field:"optional" json:"networkPerformanceConfig" yaml:"networkPerformanceConfig"`
	// pod_cidr_overprovision_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#pod_cidr_overprovision_config GoogleContainerCluster#pod_cidr_overprovision_config}
	PodCidrOverprovisionConfig *GoogleContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig `field:"optional" json:"podCidrOverprovisionConfig" yaml:"podCidrOverprovisionConfig"`
	// The IP address range for pod IPs in this node pool.
	//
	// Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#pod_ipv4_cidr_block GoogleContainerCluster#pod_ipv4_cidr_block}
	PodIpv4CidrBlock *string `field:"optional" json:"podIpv4CidrBlock" yaml:"podIpv4CidrBlock"`
	// The ID of the secondary range for pod IPs.
	//
	// If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#pod_range GoogleContainerCluster#pod_range}
	PodRange *string `field:"optional" json:"podRange" yaml:"podRange"`
	// The subnetwork name/path for the node pool.
	//
	// Format: subnetwork or projects/{project}/regions/{region}/subnetworks/{subnetwork}. This value may be specified via the nested network_config block (setting this attribute directly is supported for backward compatibility). Once created the node pool's subnetwork is immutable. If not set, the provider/API will choose the subnetwork (e.g. based on IP utilization) and report it here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#subnetwork GoogleContainerCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

