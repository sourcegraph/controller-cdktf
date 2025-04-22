package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#source GoogleOsConfigV2PolicyOrchestratorForOrganization#source}
	Source *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: 'dpkg -i package'
	// - install when true: 'apt-get update && apt-get -y install
	// package.deb'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#pull_deps GoogleOsConfigV2PolicyOrchestratorForOrganization#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

