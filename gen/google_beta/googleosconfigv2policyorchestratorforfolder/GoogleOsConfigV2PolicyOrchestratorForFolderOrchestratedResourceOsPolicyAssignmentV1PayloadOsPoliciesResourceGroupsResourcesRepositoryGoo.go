package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#name GoogleOsConfigV2PolicyOrchestratorForFolder#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#url GoogleOsConfigV2PolicyOrchestratorForFolder#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

