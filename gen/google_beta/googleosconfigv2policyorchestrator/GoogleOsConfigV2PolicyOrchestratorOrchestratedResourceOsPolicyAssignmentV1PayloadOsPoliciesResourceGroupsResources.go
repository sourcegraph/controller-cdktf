package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResources struct {
	// Required. The id of the resource with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the OS policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#id GoogleOsConfigV2PolicyOrchestrator#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#exec GoogleOsConfigV2PolicyOrchestrator#exec}
	Exec *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExec `field:"optional" json:"exec" yaml:"exec"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#file GoogleOsConfigV2PolicyOrchestrator#file}
	File *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFile `field:"optional" json:"file" yaml:"file"`
	// pkg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#pkg GoogleOsConfigV2PolicyOrchestrator#pkg}
	Pkg *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg `field:"optional" json:"pkg" yaml:"pkg"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#repository GoogleOsConfigV2PolicyOrchestrator#repository}
	Repository *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesRepository `field:"optional" json:"repository" yaml:"repository"`
}

