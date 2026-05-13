package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeGatewayIpRangeReservations struct {
	// A block of IP address ranges used to allocate supporting infrastructure for this gateway—for example, 10.1.2.0/23. The IP address block must be a /23 range. This IP address block must not overlap with subnets in any spoke or peer network that the gateway can communicate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_spoke#ip_range GoogleNetworkConnectivitySpoke#ip_range}
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
}

