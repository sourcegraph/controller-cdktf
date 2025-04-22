package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#instance_filter GoogleOsConfigV2PolicyOrchestrator#instance_filter}
	InstanceFilter *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#os_policies GoogleOsConfigV2PolicyOrchestrator#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#rollout GoogleOsConfigV2PolicyOrchestrator#rollout}
	Rollout *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#description GoogleOsConfigV2PolicyOrchestrator#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#name GoogleOsConfigV2PolicyOrchestrator#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

