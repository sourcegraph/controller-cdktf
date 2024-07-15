package googlecomputeinstance


type GoogleComputeInstanceReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#type GoogleComputeInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance#specific_reservation GoogleComputeInstance#specific_reservation}
	SpecificReservation *GoogleComputeInstanceReservationAffinitySpecificReservation `field:"optional" json:"specificReservation" yaml:"specificReservation"`
}

