package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector struct {
	// Optional. Names of the locations in scope. Format: 'us-central1-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#included_locations GoogleOsConfigV2PolicyOrchestrator#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

