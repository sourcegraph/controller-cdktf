package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyTarget struct {
	// Required when targeting forwarding rules and secure web proxy.
	//
	// Must not be specified when targeting Agent
	// Gateway. All resources referenced by this policy and extensions must share the same load balancing scheme.
	// For more information, refer to [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service). Possible values: ["INTERNAL_MANAGED", "EXTERNAL_MANAGED", "INTERNAL_SELF_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#load_balancing_scheme GoogleNetworkSecurityAuthzPolicy#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// A list of references to the Forwarding Rules or Secure Web Proxy Gateways or Agent Gateways on which this policy will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#resources GoogleNetworkSecurityAuthzPolicy#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

