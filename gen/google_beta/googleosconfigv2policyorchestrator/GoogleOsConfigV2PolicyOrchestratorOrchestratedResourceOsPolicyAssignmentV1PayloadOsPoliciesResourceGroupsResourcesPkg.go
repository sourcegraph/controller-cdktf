package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg struct {
	// Required. The desired state the agent should maintain for this package. Possible values: DESIRED_STATE_UNSPECIFIED INSTALLED REMOVED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#desired_state GoogleOsConfigV2PolicyOrchestrator#desired_state}
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#apt GoogleOsConfigV2PolicyOrchestrator#apt}
	Apt *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt `field:"optional" json:"apt" yaml:"apt"`
	// deb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#deb GoogleOsConfigV2PolicyOrchestrator#deb}
	Deb *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb `field:"optional" json:"deb" yaml:"deb"`
	// googet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#googet GoogleOsConfigV2PolicyOrchestrator#googet}
	Googet *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget `field:"optional" json:"googet" yaml:"googet"`
	// msi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#msi GoogleOsConfigV2PolicyOrchestrator#msi}
	Msi *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi `field:"optional" json:"msi" yaml:"msi"`
	// rpm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#rpm GoogleOsConfigV2PolicyOrchestrator#rpm}
	Rpm *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm `field:"optional" json:"rpm" yaml:"rpm"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#yum GoogleOsConfigV2PolicyOrchestrator#yum}
	Yum *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#zypper GoogleOsConfigV2PolicyOrchestrator#zypper}
	Zypper *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

