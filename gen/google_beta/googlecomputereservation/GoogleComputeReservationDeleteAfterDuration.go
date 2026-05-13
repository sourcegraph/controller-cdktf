package googlecomputereservation


type GoogleComputeReservationDeleteAfterDuration struct {
	// Number of nanoseconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#nanos GoogleComputeReservation#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Number of seconds for the auto-delete duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_reservation#seconds GoogleComputeReservation#seconds}
	Seconds *string `field:"optional" json:"seconds" yaml:"seconds"`
}

