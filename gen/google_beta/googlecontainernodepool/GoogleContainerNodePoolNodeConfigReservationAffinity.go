package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#consume_reservation_type GoogleContainerNodePool#consume_reservation_type}
	ConsumeReservationType *string `field:"required" json:"consumeReservationType" yaml:"consumeReservationType"`
	// The label key of a reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#key GoogleContainerNodePool#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The label values of the reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_node_pool#values GoogleContainerNodePool#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

