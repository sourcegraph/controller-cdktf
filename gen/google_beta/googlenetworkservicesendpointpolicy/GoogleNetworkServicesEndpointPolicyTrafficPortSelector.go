package googlenetworkservicesendpointpolicy


type GoogleNetworkServicesEndpointPolicyTrafficPortSelector struct {
	// List of ports.
	//
	// Can be port numbers or port range (example, [80-90] specifies all ports from 80 to 90, including 80 and 90) or named ports or * to specify all ports. If the list is empty, all ports are selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_endpoint_policy#ports GoogleNetworkServicesEndpointPolicy#ports}
	Ports *[]*string `field:"required" json:"ports" yaml:"ports"`
}

