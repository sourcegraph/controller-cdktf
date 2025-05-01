package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResource struct {
	// Optional. ID of the resource to be used while generating set of affected resources.
	//
	// For UPSERT action the value is auto-generated during PolicyOrchestrator
	// creation when not set. When the value is set it should following next
	// restrictions:
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	//
	// For DELETE action, ID must be specified explicitly during
	// PolicyOrchestrator creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#id GoogleOsConfigV2PolicyOrchestratorForOrganization#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// os_policy_assignment_v1_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#os_policy_assignment_v1_payload GoogleOsConfigV2PolicyOrchestratorForOrganization#os_policy_assignment_v1_payload}
	OsPolicyAssignmentV1Payload *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1Payload `field:"optional" json:"osPolicyAssignmentV1Payload" yaml:"osPolicyAssignmentV1Payload"`
}

