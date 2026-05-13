package googlecomputehavpngateway


type GoogleComputeHaVpnGatewayParams struct {
	// Resource manager tags to be bound to the HaVpnGateway.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_ha_vpn_gateway#resource_manager_tags GoogleComputeHaVpnGateway#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

