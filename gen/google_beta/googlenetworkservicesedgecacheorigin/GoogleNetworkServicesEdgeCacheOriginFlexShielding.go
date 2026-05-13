package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginFlexShielding struct {
	// Whenever possible, content will be fetched from origin and cached in or near the specified origin. Best effort.
	//
	// You must specify exactly one FlexShieldingRegion. Possible values: ["AFRICA_SOUTH1", "ME_CENTRAL1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_edge_cache_origin#flex_shielding_regions GoogleNetworkServicesEdgeCacheOrigin#flex_shielding_regions}
	FlexShieldingRegions *[]*string `field:"optional" json:"flexShieldingRegions" yaml:"flexShieldingRegions"`
}

