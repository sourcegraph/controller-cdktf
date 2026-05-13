package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#location_selector GoogleOsConfigV2PolicyOrchestratorForFolder#location_selector}
	LocationSelector *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#resource_hierarchy_selector GoogleOsConfigV2PolicyOrchestratorForFolder#resource_hierarchy_selector}
	ResourceHierarchySelector *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

