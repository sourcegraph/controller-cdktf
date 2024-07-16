package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicySignedTokenOptions struct {
	// The allowed signature algorithms to use.
	//
	// Defaults to using only ED25519.
	//
	// You may specify up to 3 signature algorithms to use. Possible values: ["ED25519", "HMAC_SHA_256", "HMAC_SHA1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#allowed_signature_algorithms GoogleNetworkServicesEdgeCacheService#allowed_signature_algorithms}
	AllowedSignatureAlgorithms *[]*string `field:"optional" json:"allowedSignatureAlgorithms" yaml:"allowedSignatureAlgorithms"`
	// The query parameter in which to find the token.
	//
	// The name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.
	//
	// Defaults to 'edge-cache-token'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#token_query_parameter GoogleNetworkServicesEdgeCacheService#token_query_parameter}
	TokenQueryParameter *string `field:"optional" json:"tokenQueryParameter" yaml:"tokenQueryParameter"`
}

