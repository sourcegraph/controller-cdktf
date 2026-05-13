package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsLocationSelector struct {
	// Names of the locations in scope. Format: 'us-central1-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#included_locations GoogleOsConfigV2PolicyOrchestratorForFolder#included_locations}
	IncludedLocations *[]*string `field:"optional" json:"includedLocations" yaml:"includedLocations"`
}

