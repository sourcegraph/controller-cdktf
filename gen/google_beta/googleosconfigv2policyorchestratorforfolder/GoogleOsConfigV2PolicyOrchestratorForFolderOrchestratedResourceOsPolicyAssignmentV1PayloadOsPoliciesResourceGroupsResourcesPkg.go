package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg struct {
	// The desired state the agent should maintain for this package. Possible values: ["INSTALLED", "REMOVED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#desired_state GoogleOsConfigV2PolicyOrchestratorForFolder#desired_state}
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#apt GoogleOsConfigV2PolicyOrchestratorForFolder#apt}
	Apt *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt `field:"optional" json:"apt" yaml:"apt"`
	// deb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#deb GoogleOsConfigV2PolicyOrchestratorForFolder#deb}
	Deb *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb `field:"optional" json:"deb" yaml:"deb"`
	// googet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#googet GoogleOsConfigV2PolicyOrchestratorForFolder#googet}
	Googet *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget `field:"optional" json:"googet" yaml:"googet"`
	// msi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#msi GoogleOsConfigV2PolicyOrchestratorForFolder#msi}
	Msi *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi `field:"optional" json:"msi" yaml:"msi"`
	// rpm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#rpm GoogleOsConfigV2PolicyOrchestratorForFolder#rpm}
	Rpm *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm `field:"optional" json:"rpm" yaml:"rpm"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#yum GoogleOsConfigV2PolicyOrchestratorForFolder#yum}
	Yum *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#zypper GoogleOsConfigV2PolicyOrchestratorForFolder#zypper}
	Zypper *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

