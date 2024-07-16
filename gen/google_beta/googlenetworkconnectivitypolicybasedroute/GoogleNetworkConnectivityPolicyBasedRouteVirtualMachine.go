package googlenetworkconnectivitypolicybasedroute


type GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine struct {
	// A list of VM instance tags that this policy-based route applies to.
	//
	// VM instances that have ANY of tags specified here will install this PBR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#tags GoogleNetworkConnectivityPolicyBasedRoute#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
}

