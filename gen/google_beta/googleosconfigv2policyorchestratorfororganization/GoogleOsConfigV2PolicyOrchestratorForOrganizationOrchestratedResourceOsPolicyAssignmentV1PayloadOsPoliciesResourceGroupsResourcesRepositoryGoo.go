package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// Required. The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#name GoogleOsConfigV2PolicyOrchestratorForOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#url GoogleOsConfigV2PolicyOrchestratorForOrganization#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

