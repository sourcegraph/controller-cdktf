package googlecomputewiregroup


type GoogleComputeWireGroupWireGroupProperties struct {
	// Type of wire group (enum).
	//
	// WIRE: a single pseudowire over two Interconnect connections   with no redundancy.
	// REDUNDANT: two pseudowires over four Interconnect connections, with two connections in one metro and two connections in another metro.
	// BOX_AND_CROSS: four pseudowires over four Interconnect connections, with two connections in one metro and two connections in another metro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#type GoogleComputeWireGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

