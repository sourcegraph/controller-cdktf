package googlecomputerolloutplan


type GoogleComputeRolloutPlanWavesValidation struct {
	// The type of the validation.
	//
	// Possible values:
	// "manual": The system waits for an end-user approval API before progressing to the next wave.
	// "time": The system waits for a user specified duration before progressing to the next wave.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#type GoogleComputeRolloutPlan#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// time_based_validation_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#time_based_validation_metadata GoogleComputeRolloutPlan#time_based_validation_metadata}
	TimeBasedValidationMetadata *GoogleComputeRolloutPlanWavesValidationTimeBasedValidationMetadata `field:"optional" json:"timeBasedValidationMetadata" yaml:"timeBasedValidationMetadata"`
}

