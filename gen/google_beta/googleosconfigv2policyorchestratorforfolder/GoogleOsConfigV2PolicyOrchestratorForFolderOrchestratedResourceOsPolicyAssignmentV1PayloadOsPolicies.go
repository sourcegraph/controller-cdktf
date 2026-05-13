package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPolicies struct {
	// The id of the OS policy with the following restrictions:.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#id GoogleOsConfigV2PolicyOrchestratorForFolder#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Policy mode Possible values: ["VALIDATION", "ENFORCEMENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#mode GoogleOsConfigV2PolicyOrchestratorForFolder#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// resource_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#resource_groups GoogleOsConfigV2PolicyOrchestratorForFolder#resource_groups}
	ResourceGroups interface{} `field:"required" json:"resourceGroups" yaml:"resourceGroups"`
	// This flag determines the OS policy compliance status when none of the resource groups within the policy are applicable for a VM.
	//
	// Set this value
	// to 'true' if the policy needs to be reported as compliant even if the
	// policy has nothing to validate or enforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#allow_no_resource_group_match GoogleOsConfigV2PolicyOrchestratorForFolder#allow_no_resource_group_match}
	AllowNoResourceGroupMatch interface{} `field:"optional" json:"allowNoResourceGroupMatch" yaml:"allowNoResourceGroupMatch"`
	// Policy description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#description GoogleOsConfigV2PolicyOrchestratorForFolder#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

