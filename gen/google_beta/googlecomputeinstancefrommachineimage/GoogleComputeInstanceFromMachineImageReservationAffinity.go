package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#type GoogleComputeInstanceFromMachineImage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_machine_image#specific_reservation GoogleComputeInstanceFromMachineImage#specific_reservation}
	SpecificReservation *GoogleComputeInstanceFromMachineImageReservationAffinitySpecificReservation `field:"optional" json:"specificReservation" yaml:"specificReservation"`
}

