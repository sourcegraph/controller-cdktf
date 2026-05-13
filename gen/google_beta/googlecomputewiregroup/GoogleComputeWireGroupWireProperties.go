package googlecomputewiregroup


type GoogleComputeWireGroupWireProperties struct {
	// The configuration of a wire's bandwidth allocation.
	//
	// ALLOCATE_PER_WIRE: configures a separate unmetered bandwidth allocation (and associated charges) for each wire in the group.
	// SHARED_WITH_WIRE_GROUP: this is the default behavior, which configures one unmetered bandwidth allocation for the wire group. The unmetered bandwidth is divided equally across each wire in the group, but dynamic
	// throttling reallocates unused unmetered bandwidth from unused or underused wires to other wires in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#bandwidth_allocation GoogleComputeWireGroup#bandwidth_allocation}
	BandwidthAllocation *string `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// The unmetered bandwidth setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#bandwidth_unmetered GoogleComputeWireGroup#bandwidth_unmetered}
	BandwidthUnmetered *float64 `field:"optional" json:"bandwidthUnmetered" yaml:"bandwidthUnmetered"`
	// Response when a fault is detected in a pseudowire: NONE: default.
	//
	// DISABLE_PORT: set the port line protocol down when inline probes detect a fault. This setting is only permitted on port mode pseudowires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_wire_group#fault_response GoogleComputeWireGroup#fault_response}
	FaultResponse *string `field:"optional" json:"faultResponse" yaml:"faultResponse"`
}

