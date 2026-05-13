package googlecomputefuturereservation


type GoogleComputeFutureReservationCommitmentInfo struct {
	// name of the commitment where capacity is being delivered to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#commitment_name GoogleComputeFutureReservation#commitment_name}
	CommitmentName *string `field:"optional" json:"commitmentName" yaml:"commitmentName"`
	// Indicates if a Commitment needs to be created as part of FR delivery.
	//
	// If this field is not present, then no commitment needs to be created. Possible values: ["INVALID", "THIRTY_SIX_MONTH", "TWELVE_MONTH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#commitment_plan GoogleComputeFutureReservation#commitment_plan}
	CommitmentPlan *string `field:"optional" json:"commitmentPlan" yaml:"commitmentPlan"`
	// Only applicable if FR is delivering to the same reservation.
	//
	// If set, all parent commitments will be extended to match the end date of the plan for this commitment. Possible values: ["EXTEND"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#previous_commitment_terms GoogleComputeFutureReservation#previous_commitment_terms}
	PreviousCommitmentTerms *string `field:"optional" json:"previousCommitmentTerms" yaml:"previousCommitmentTerms"`
}

