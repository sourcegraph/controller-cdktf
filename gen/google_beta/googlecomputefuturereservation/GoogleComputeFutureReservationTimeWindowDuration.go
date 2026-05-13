package googlecomputefuturereservation


type GoogleComputeFutureReservationTimeWindowDuration struct {
	// Span of time that's a fraction of a second at nanosecond resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#nanos GoogleComputeFutureReservation#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#seconds GoogleComputeFutureReservation#seconds}
	Seconds *string `field:"optional" json:"seconds" yaml:"seconds"`
}

