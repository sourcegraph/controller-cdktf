package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#source GoogleOsConfigV2PolicyOrchestratorForOrganization#source}
	Source *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsiSource `field:"required" json:"source" yaml:"source"`
	// Additional properties to use during installation. This should be in the format of Property=Setting. Appended to the defaults of 'ACTION=INSTALL REBOOT=ReallySuppress'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#properties GoogleOsConfigV2PolicyOrchestratorForOrganization#properties}
	Properties *[]*string `field:"optional" json:"properties" yaml:"properties"`
}

