package computewiregroup


type ComputeWireGroupEndpointsInterconnects struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_wire_group#interconnect_name ComputeWireGroup#interconnect_name}.
	InterconnectName *string `field:"required" json:"interconnectName" yaml:"interconnectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_wire_group#interconnect ComputeWireGroup#interconnect}.
	Interconnect *string `field:"optional" json:"interconnect" yaml:"interconnect"`
	// VLAN tags for the interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_wire_group#vlan_tags ComputeWireGroup#vlan_tags}
	VlanTags *[]*float64 `field:"optional" json:"vlanTags" yaml:"vlanTags"`
}

