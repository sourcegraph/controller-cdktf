package googlecomputeinterconnectgroup


type GoogleComputeInterconnectGroupInterconnects struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_group#name GoogleComputeInterconnectGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of an Interconnect in this group. All Interconnects in the group are unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_interconnect_group#interconnect GoogleComputeInterconnectGroup#interconnect}
	Interconnect *string `field:"optional" json:"interconnect" yaml:"interconnect"`
}

