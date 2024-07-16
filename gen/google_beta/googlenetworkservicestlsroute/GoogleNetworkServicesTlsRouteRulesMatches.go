package googlenetworkservicestlsroute


type GoogleNetworkServicesTlsRouteRulesMatches struct {
	// ALPN (Application-Layer Protocol Negotiation) to match against.
	//
	// Examples: "http/1.1", "h2". At least one of sniHost and alpn is required. Up to 5 alpns across all matches can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#alpn GoogleNetworkServicesTlsRoute#alpn}
	Alpn *[]*string `field:"optional" json:"alpn" yaml:"alpn"`
	// SNI (server name indicator) to match against.
	//
	// SNI will be matched against all wildcard domains, i.e. www.example.com will be first matched against www.example.com, then *.example.com, then *.com.
	// Partial wildcards are not supported, and values like *w.example.com are invalid. At least one of sniHost and alpn is required. Up to 5 sni hosts across all matches can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#sni_host GoogleNetworkServicesTlsRoute#sni_host}
	SniHost *[]*string `field:"optional" json:"sniHost" yaml:"sniHost"`
}

