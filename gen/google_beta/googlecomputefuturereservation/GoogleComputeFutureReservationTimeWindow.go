package googlecomputefuturereservation


type GoogleComputeFutureReservationTimeWindow struct {
	// Start time of the future reservation in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#start_time GoogleComputeFutureReservation#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#duration GoogleComputeFutureReservation#duration}
	Duration *GoogleComputeFutureReservationTimeWindowDuration `field:"optional" json:"duration" yaml:"duration"`
	// End time of the future reservation in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#end_time GoogleComputeFutureReservation#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
}

