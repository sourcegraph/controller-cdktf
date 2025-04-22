package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#location_selector GoogleOsConfigV2PolicyOrchestrator#location_selector}
	LocationSelector *GoogleOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#resource_hierarchy_selector GoogleOsConfigV2PolicyOrchestrator#resource_hierarchy_selector}
	ResourceHierarchySelector *GoogleOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

