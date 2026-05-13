package lustreinstance


type LustreInstanceAccessRulesOptionsAccessRules struct {
	// The IP address ranges to which to apply this access rule.
	//
	// Accepts
	// non-overlapping CIDR ranges (e.g., '192.168.1.0/24') and IP addresses
	// (e.g., '192.168.1.0').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#ip_address_ranges LustreInstance#ip_address_ranges}
	IpAddressRanges *[]*string `field:"required" json:"ipAddressRanges" yaml:"ipAddressRanges"`
	// The name of the access rule policy group. Must be 16 characters or less and include only alphanumeric characters or '_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#name LustreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Squash mode for the access rule. Possible values: NO_SQUASH ROOT_SQUASH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#squash_mode LustreInstance#squash_mode}
	SquashMode *string `field:"required" json:"squashMode" yaml:"squashMode"`
}

