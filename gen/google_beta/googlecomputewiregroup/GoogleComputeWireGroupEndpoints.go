package googlecomputewiregroup


type GoogleComputeWireGroupEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#endpoint GoogleComputeWireGroup#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// interconnects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#interconnects GoogleComputeWireGroup#interconnects}
	Interconnects interface{} `field:"optional" json:"interconnects" yaml:"interconnects"`
}

