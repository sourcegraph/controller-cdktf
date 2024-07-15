package googlenetworkconnectivityserviceconnectionpolicy


type GoogleNetworkConnectivityServiceConnectionPolicyPscConfig struct {
	// IDs of the subnetworks or fully qualified identifiers for the subnetworks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_service_connection_policy#subnetworks GoogleNetworkConnectivityServiceConnectionPolicy#subnetworks}
	Subnetworks *[]*string `field:"required" json:"subnetworks" yaml:"subnetworks"`
	// Max number of PSC connections for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_service_connection_policy#limit GoogleNetworkConnectivityServiceConnectionPolicy#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
}

