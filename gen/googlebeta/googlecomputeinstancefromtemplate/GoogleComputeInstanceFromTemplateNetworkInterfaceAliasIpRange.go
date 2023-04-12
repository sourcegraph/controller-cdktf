package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateNetworkInterfaceAliasIpRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#ip_cidr_range GoogleComputeInstanceFromTemplate#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#subnetwork_range_name GoogleComputeInstanceFromTemplate#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"optional" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

