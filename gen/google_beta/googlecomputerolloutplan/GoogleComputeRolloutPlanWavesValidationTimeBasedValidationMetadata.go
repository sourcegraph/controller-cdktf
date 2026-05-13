package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesValidationTimeBasedValidationMetadata struct {
	// The duration that the system waits in between waves.
	//
	// This wait starts
	// after all changes in the wave are rolled out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#wait_duration GoogleComputeRolloutPlan#wait_duration}
	WaitDuration *string `field:"optional" json:"waitDuration" yaml:"waitDuration"`
}

