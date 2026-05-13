package networkconnectivitydestination


type NetworkConnectivityDestinationEndpoints struct {
	// The ASN of the remote IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#asn NetworkConnectivityDestination#asn}
	Asn *string `field:"required" json:"asn" yaml:"asn"`
	// The CSP of the remote IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#csp NetworkConnectivityDestination#csp}
	Csp *string `field:"required" json:"csp" yaml:"csp"`
}

