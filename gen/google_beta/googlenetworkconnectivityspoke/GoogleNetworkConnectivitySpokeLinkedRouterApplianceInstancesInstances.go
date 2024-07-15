package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances struct {
	// The IP address on the VM to use for peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_spoke#ip_address GoogleNetworkConnectivitySpoke#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The URI of the virtual machine resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_connectivity_spoke#virtual_machine GoogleNetworkConnectivitySpoke#virtual_machine}
	VirtualMachine *string `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

