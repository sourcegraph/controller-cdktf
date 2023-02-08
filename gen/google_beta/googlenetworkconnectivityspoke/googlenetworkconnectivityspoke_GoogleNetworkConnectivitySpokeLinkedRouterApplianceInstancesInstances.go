package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances struct {
	// The IP address on the VM to use for peering.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_connectivity_spoke#ip_address GoogleNetworkConnectivitySpoke#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The URI of the virtual machine resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_network_connectivity_spoke#virtual_machine GoogleNetworkConnectivitySpoke#virtual_machine}
	VirtualMachine *string `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

