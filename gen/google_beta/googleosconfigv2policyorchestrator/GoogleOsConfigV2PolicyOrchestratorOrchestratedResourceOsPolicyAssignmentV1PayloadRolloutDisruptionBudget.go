package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#fixed GoogleOsConfigV2PolicyOrchestrator#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#percent GoogleOsConfigV2PolicyOrchestrator#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

