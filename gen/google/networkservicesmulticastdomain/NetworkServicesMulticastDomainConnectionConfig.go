package networkservicesmulticastdomain


type NetworkServicesMulticastDomainConnectionConfig struct {
	// The VPC connection type. Possible values: NCC SAME_VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain#connection_type NetworkServicesMulticastDomain#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// The resource name of the [NCC](https://cloud.google.com/network-connectivity-center) hub. Use the following format: 'projects/{project}/locations/global/hubs/{hub}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_domain#ncc_hub NetworkServicesMulticastDomain#ncc_hub}
	NccHub *string `field:"optional" json:"nccHub" yaml:"nccHub"`
}

