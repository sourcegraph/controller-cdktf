package googlenetworkconnectivitydestination


type GoogleNetworkConnectivityDestinationEndpoints struct {
	// The ASN of the remote IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_destination#asn GoogleNetworkConnectivityDestination#asn}
	Asn *string `field:"required" json:"asn" yaml:"asn"`
	// The CSP of the remote IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_destination#csp GoogleNetworkConnectivityDestination#csp}
	Csp *string `field:"required" json:"csp" yaml:"csp"`
}

