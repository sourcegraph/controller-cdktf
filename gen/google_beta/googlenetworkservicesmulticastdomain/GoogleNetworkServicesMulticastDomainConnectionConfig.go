package googlenetworkservicesmulticastdomain


type GoogleNetworkServicesMulticastDomainConnectionConfig struct {
	// The VPC connection type. Possible values: NCC SAME_VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#connection_type GoogleNetworkServicesMulticastDomain#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// The resource name of the [NCC](https://cloud.google.com/network-connectivity-center) hub. Use the following format: 'projects/{project}/locations/global/hubs/{hub}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain#ncc_hub GoogleNetworkServicesMulticastDomain#ncc_hub}
	NccHub *string `field:"optional" json:"nccHub" yaml:"nccHub"`
}

