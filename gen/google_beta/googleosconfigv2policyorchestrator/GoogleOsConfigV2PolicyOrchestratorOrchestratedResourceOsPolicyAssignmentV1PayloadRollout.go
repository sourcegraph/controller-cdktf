package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#disruption_budget GoogleOsConfigV2PolicyOrchestrator#disruption_budget}
	DisruptionBudget *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// Required.
	//
	// This determines the minimum duration of time to wait after the
	// configuration changes are applied through the current rollout. A
	// VM continues to count towards the 'disruption_budget' at least
	// until this duration of time has passed after configuration changes are
	// applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#min_wait_duration GoogleOsConfigV2PolicyOrchestrator#min_wait_duration}
	MinWaitDuration *string `field:"required" json:"minWaitDuration" yaml:"minWaitDuration"`
}

