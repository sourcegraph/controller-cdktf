package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterNetworkConfig struct {
	// Enables the use of advanced Anthos networking features, such as Bundled Load Balancing with BGP or the egress NAT gateway.
	//
	// Setting configuration for advanced networking features will automatically
	// set this flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#advanced_networking GoogleGkeonpremBareMetalCluster#advanced_networking}
	AdvancedNetworking interface{} `field:"optional" json:"advancedNetworking" yaml:"advancedNetworking"`
	// island_mode_cidr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#island_mode_cidr GoogleGkeonpremBareMetalCluster#island_mode_cidr}
	IslandModeCidr *GoogleGkeonpremBareMetalClusterNetworkConfigIslandModeCidr `field:"optional" json:"islandModeCidr" yaml:"islandModeCidr"`
	// multiple_network_interfaces_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#multiple_network_interfaces_config GoogleGkeonpremBareMetalCluster#multiple_network_interfaces_config}
	MultipleNetworkInterfacesConfig *GoogleGkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig `field:"optional" json:"multipleNetworkInterfacesConfig" yaml:"multipleNetworkInterfacesConfig"`
	// sr_iov_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#sr_iov_config GoogleGkeonpremBareMetalCluster#sr_iov_config}
	SrIovConfig *GoogleGkeonpremBareMetalClusterNetworkConfigSrIovConfig `field:"optional" json:"srIovConfig" yaml:"srIovConfig"`
}

