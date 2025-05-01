package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryApt struct {
	// Required. Type of archive files in this repository. Possible values: ARCHIVE_TYPE_UNSPECIFIED DEB DEB_SRC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#archive_type GoogleOsConfigV2PolicyOrchestratorForOrganization#archive_type}
	ArchiveType *string `field:"required" json:"archiveType" yaml:"archiveType"`
	// Required. List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#components GoogleOsConfigV2PolicyOrchestratorForOrganization#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Required. Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#distribution GoogleOsConfigV2PolicyOrchestratorForOrganization#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// Required. URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#uri GoogleOsConfigV2PolicyOrchestratorForOrganization#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// URI of the key file for this repository. The agent maintains a keyring at '/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#gpg_key GoogleOsConfigV2PolicyOrchestratorForOrganization#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

