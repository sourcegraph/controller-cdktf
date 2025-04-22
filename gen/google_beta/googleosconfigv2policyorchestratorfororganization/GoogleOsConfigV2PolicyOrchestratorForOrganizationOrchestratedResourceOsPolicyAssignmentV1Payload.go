package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1Payload struct {
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#instance_filter GoogleOsConfigV2PolicyOrchestratorForOrganization#instance_filter}
	InstanceFilter *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#os_policies GoogleOsConfigV2PolicyOrchestratorForOrganization#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#rollout GoogleOsConfigV2PolicyOrchestratorForOrganization#rollout}
	Rollout *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#description GoogleOsConfigV2PolicyOrchestratorForOrganization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#etag GoogleOsConfigV2PolicyOrchestratorForOrganization#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Resource name.
	//
	// Format:
	// 'projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}'
	//
	// This field is ignored when you create an OS policy assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#name GoogleOsConfigV2PolicyOrchestratorForOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

