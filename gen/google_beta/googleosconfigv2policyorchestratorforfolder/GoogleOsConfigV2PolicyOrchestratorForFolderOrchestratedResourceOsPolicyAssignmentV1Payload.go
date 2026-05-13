package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#instance_filter GoogleOsConfigV2PolicyOrchestratorForFolder#instance_filter}
	InstanceFilter *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#os_policies GoogleOsConfigV2PolicyOrchestratorForFolder#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#rollout GoogleOsConfigV2PolicyOrchestratorForFolder#rollout}
	Rollout *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#description GoogleOsConfigV2PolicyOrchestratorForFolder#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#name GoogleOsConfigV2PolicyOrchestratorForFolder#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

