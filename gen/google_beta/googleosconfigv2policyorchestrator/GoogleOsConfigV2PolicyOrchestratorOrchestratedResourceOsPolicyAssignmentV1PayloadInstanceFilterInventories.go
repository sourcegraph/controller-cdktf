package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilterInventories struct {
	// Required. The OS short name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#os_short_name GoogleOsConfigV2PolicyOrchestrator#os_short_name}
	OsShortName *string `field:"required" json:"osShortName" yaml:"osShortName"`
	// The OS version.
	//
	// Prefix matches are supported if asterisk(*) is provided as the
	// last character. For example, to match all versions with a major
	// version of '7', specify the following value for this field '7.*'
	//
	// An empty string matches all OS versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_os_config_v2_policy_orchestrator#os_version GoogleOsConfigV2PolicyOrchestrator#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
}

