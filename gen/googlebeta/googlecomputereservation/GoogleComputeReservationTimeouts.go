package googlecomputereservation


type GoogleComputeReservationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_reservation#create GoogleComputeReservation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_reservation#delete GoogleComputeReservation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_reservation#update GoogleComputeReservation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

