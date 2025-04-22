package googlecomputesubnetwork


type GoogleComputeSubnetworkSecondaryIpRange struct {
	// The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance.
	//
	// The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_subnetwork#range_name GoogleComputeSubnetwork#range_name}
	RangeName *string `field:"required" json:"rangeName" yaml:"rangeName"`
	// The range of IP addresses belonging to this subnetwork secondary range.
	//
	// Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	// Field is optional when 'reserved_internal_range' is defined, otherwise required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g. 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_subnetwork#reserved_internal_range GoogleComputeSubnetwork#reserved_internal_range}
	ReservedInternalRange *string `field:"optional" json:"reservedInternalRange" yaml:"reservedInternalRange"`
}

