package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesFileFileGcs struct {
	// Required. Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#bucket GoogleOsConfigV2PolicyOrchestrator#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Required. Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#object GoogleOsConfigV2PolicyOrchestrator#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#generation GoogleOsConfigV2PolicyOrchestrator#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

