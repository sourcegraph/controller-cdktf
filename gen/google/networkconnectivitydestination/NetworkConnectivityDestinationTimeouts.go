package networkconnectivitydestination


type NetworkConnectivityDestinationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#create NetworkConnectivityDestination#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#delete NetworkConnectivityDestination#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#update NetworkConnectivityDestination#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

