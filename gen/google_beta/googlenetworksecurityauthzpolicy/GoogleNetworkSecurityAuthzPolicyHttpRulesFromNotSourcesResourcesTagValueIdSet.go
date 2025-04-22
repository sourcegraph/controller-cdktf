package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesResourcesTagValueIdSet struct {
	// A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request.
	//
	// The match follows AND semantics which means all the ids must match.
	// Limited to 5 matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#ids GoogleNetworkSecurityAuthzPolicy#ids}
	Ids *[]*string `field:"optional" json:"ids" yaml:"ids"`
}

