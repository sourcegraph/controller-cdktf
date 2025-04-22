package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateNetworkInterfaceAliasIpRange struct {
	// The IP CIDR range represented by this alias IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#ip_cidr_range GoogleComputeInstanceFromTemplate#ip_cidr_range}
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#subnetwork_range_name GoogleComputeInstanceFromTemplate#subnetwork_range_name}
	SubnetworkRangeName *string `field:"optional" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

