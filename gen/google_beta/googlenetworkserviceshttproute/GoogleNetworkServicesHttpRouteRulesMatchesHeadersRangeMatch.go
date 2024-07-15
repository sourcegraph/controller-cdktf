package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesMatchesHeadersRangeMatch struct {
	// End of the range (exclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#end GoogleNetworkServicesHttpRoute#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Start of the range (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#start GoogleNetworkServicesHttpRoute#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

