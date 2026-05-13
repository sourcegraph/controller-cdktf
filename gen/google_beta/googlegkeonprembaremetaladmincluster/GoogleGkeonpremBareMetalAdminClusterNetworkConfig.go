package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterNetworkConfig struct {
	// Enables the use of advanced Anthos networking features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#advanced_networking GoogleGkeonpremBareMetalAdminCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// island_mode_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#island_mode_cidr GoogleGkeonpremBareMetalAdminCluster#island_mode_cidr}
	IslandModeCidr *GoogleGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr `field:"optional" json:"islandModeCidr" yaml:"islandModeCidr"`
	// multiple_network_interfaces_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#multiple_network_interfaces_config GoogleGkeonpremBareMetalAdminCluster#multiple_network_interfaces_config}
	MultipleNetworkInterfacesConfig *GoogleGkeonpremBareMetalAdminClusterNetworkConfigMultipleNetworkInterfacesConfig `field:"optional" json:"multipleNetworkInterfacesConfig" yaml:"multipleNetworkInterfacesConfig"`
}

