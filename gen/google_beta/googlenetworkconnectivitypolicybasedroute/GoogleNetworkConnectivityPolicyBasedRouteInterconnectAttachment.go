package googlenetworkconnectivitypolicybasedroute


type GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment struct {
	// Cloud region to install this policy-based route on for Interconnect attachments.
	//
	// Use 'all' to install it on all Interconnect attachments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#region GoogleNetworkConnectivityPolicyBasedRoute#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

