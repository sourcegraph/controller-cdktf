package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeGateway struct {
	// the capacity of the gateway spoke, in Gbps. Possible values: ["CAPACITY_1_GBPS", "CAPACITY_10_GBPS", "CAPACITY_100_GBPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_spoke#capacity GoogleNetworkConnectivitySpoke#capacity}
	Capacity *string `field:"required" json:"capacity" yaml:"capacity"`
	// ip_range_reservations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_spoke#ip_range_reservations GoogleNetworkConnectivitySpoke#ip_range_reservations}
	IpRangeReservations interface{} `field:"required" json:"ipRangeReservations" yaml:"ipRangeReservations"`
}

