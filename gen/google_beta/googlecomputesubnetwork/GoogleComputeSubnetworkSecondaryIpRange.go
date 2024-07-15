package googlecomputesubnetwork


type GoogleComputeSubnetworkSecondaryIpRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_subnetwork#range_name GoogleComputeSubnetwork#range_name}.
	RangeName *string `field:"optional" json:"rangeName" yaml:"rangeName"`
}

