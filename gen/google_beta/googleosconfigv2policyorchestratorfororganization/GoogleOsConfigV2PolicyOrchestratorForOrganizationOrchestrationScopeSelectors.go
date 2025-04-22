package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectors struct {
	// location_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#location_selector GoogleOsConfigV2PolicyOrchestratorForOrganization#location_selector}
	LocationSelector *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsLocationSelector `field:"optional" json:"locationSelector" yaml:"locationSelector"`
	// resource_hierarchy_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#resource_hierarchy_selector GoogleOsConfigV2PolicyOrchestratorForOrganization#resource_hierarchy_selector}
	ResourceHierarchySelector *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationScopeSelectorsResourceHierarchySelector `field:"optional" json:"resourceHierarchySelector" yaml:"resourceHierarchySelector"`
}

