package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#source GoogleOsConfigV2PolicyOrchestrator#source}
	Source *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: 'rpm --upgrade --replacepkgs package.rpm'
	// - install when true: 'yum -y install package.rpm' or
	// 'zypper -y install package.rpm'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#pull_deps GoogleOsConfigV2PolicyOrchestrator#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

