package googlenotebooksinstance


type GoogleNotebooksInstanceReservationAffinity struct {
	// The type of Compute Reservation. Possible values: ["NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#consume_reservation_type GoogleNotebooksInstance#consume_reservation_type}
	ConsumeReservationType *string `field:"required" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#key GoogleNotebooksInstance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance#values GoogleNotebooksInstance#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

