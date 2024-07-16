package googlenetworkconnectivitypolicybasedroute


type GoogleNetworkConnectivityPolicyBasedRouteFilter struct {
	// Internet protocol versions this policy-based route applies to. Possible values: ["IPV4"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#protocol_version GoogleNetworkConnectivityPolicyBasedRoute#protocol_version}
	ProtocolVersion *string `field:"required" json:"protocolVersion" yaml:"protocolVersion"`
	// The destination IP range of outgoing packets that this policy-based route applies to.
	//
	// Default is "0.0.0.0/0" if protocol version is IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#dest_range GoogleNetworkConnectivityPolicyBasedRoute#dest_range}
	DestRange *string `field:"optional" json:"destRange" yaml:"destRange"`
	// The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#ip_protocol GoogleNetworkConnectivityPolicyBasedRoute#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The source IP range of outgoing packets that this policy-based route applies to.
	//
	// Default is "0.0.0.0/0" if protocol version is IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_policy_based_route#src_range GoogleNetworkConnectivityPolicyBasedRoute#src_range}
	SrcRange *string `field:"optional" json:"srcRange" yaml:"srcRange"`
}

