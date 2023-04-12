package googlecomputesubnetwork


type GoogleComputeSubnetworkSecondaryIpRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork#range_name GoogleComputeSubnetwork#range_name}.
	RangeName *string `field:"optional" json:"rangeName" yaml:"rangeName"`
}

