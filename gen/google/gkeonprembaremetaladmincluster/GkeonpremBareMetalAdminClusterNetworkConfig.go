package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterNetworkConfig struct {
	// Enables the use of advanced Anthos networking features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#advanced_networking GkeonpremBareMetalAdminCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// island_mode_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#island_mode_cidr GkeonpremBareMetalAdminCluster#island_mode_cidr}
	IslandModeCidr *GkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr `field:"optional" json:"islandModeCidr" yaml:"islandModeCidr"`
	// multiple_network_interfaces_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/gkeonprem_bare_metal_admin_cluster#multiple_network_interfaces_config GkeonpremBareMetalAdminCluster#multiple_network_interfaces_config}
	MultipleNetworkInterfacesConfig *GkeonpremBareMetalAdminClusterNetworkConfigMultipleNetworkInterfacesConfig `field:"optional" json:"multipleNetworkInterfacesConfig" yaml:"multipleNetworkInterfacesConfig"`
}

