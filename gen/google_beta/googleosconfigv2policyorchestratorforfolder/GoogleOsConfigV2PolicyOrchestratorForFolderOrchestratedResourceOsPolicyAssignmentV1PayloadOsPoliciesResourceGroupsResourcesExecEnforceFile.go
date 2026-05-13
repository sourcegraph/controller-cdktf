package googleosconfigv2policyorchestratorforfolder


type GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforceFile struct {
	// Defaults to false. When false, files are subject to validations based on the file type:.
	//
	// Remote: A checksum must be specified.
	// Cloud Storage: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#allow_insecure GoogleOsConfigV2PolicyOrchestratorForFolder#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#gcs GoogleOsConfigV2PolicyOrchestratorForFolder#gcs}
	Gcs *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforceFileGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// A local path within the VM to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#local_path GoogleOsConfigV2PolicyOrchestratorForFolder#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_os_config_v2_policy_orchestrator_for_folder#remote GoogleOsConfigV2PolicyOrchestratorForFolder#remote}
	Remote *GoogleOsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecEnforceFileRemote `field:"optional" json:"remote" yaml:"remote"`
}

