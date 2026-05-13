package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Type of archive files in this repository. Possible values: ["DEB", "DEB_SRC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#archive_type GoogleOsConfigV2PolicyOrchestratorForFolder#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#components GoogleOsConfigV2PolicyOrchestratorForFolder#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#distribution GoogleOsConfigV2PolicyOrchestratorForFolder#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#uri GoogleOsConfigV2PolicyOrchestratorForFolder#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#gpg_key GoogleOsConfigV2PolicyOrchestratorForFolder#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

