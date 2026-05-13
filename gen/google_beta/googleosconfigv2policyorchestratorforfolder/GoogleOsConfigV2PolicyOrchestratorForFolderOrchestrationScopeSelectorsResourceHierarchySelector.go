package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestrationScopeSelectorsResourceHierarchySelector struct {
	// Names of the folders in scope. Format: 'folders/{folder_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#included_folders GoogleOsConfigV2PolicyOrchestratorForFolder#included_folders}
	IncludedFolders *[]*string `field:"optional" json:"includedFolders" yaml:"includedFolders"`
	// Names of the projects in scope. Format: 'projects/{project_number}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#included_projects GoogleOsConfigV2PolicyOrchestratorForFolder#included_projects}
	IncludedProjects *[]*string `field:"optional" json:"includedProjects" yaml:"includedProjects"`
}

