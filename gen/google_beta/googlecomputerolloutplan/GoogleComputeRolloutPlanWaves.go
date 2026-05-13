package googlecomputerolloutplan


type GoogleComputeRolloutPlanWaves struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#selectors GoogleComputeRolloutPlan#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#validation GoogleComputeRolloutPlan#validation}
	Validation *GoogleComputeRolloutPlanWavesValidation `field:"required" json:"validation" yaml:"validation"`
	// The display name of this wave of the rollout plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#display_name GoogleComputeRolloutPlan#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// orchestration_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_rollout_plan#orchestration_options GoogleComputeRolloutPlan#orchestration_options}
	OrchestrationOptions *GoogleComputeRolloutPlanWavesOrchestrationOptions `field:"optional" json:"orchestrationOptions" yaml:"orchestrationOptions"`
}

