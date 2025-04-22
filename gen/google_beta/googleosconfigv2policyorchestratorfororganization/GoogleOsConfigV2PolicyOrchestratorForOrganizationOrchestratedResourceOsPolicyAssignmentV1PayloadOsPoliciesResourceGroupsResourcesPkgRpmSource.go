package googleosconfigv2policyorchestratorfororganization


type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSource struct {
	// Defaults to false. When false, files are subject to validations based on the file type:.
	//
	// Remote: A checksum must be specified.
	// Cloud Storage: An object generation number must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#allow_insecure GoogleOsConfigV2PolicyOrchestratorForOrganization#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// gcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#gcs GoogleOsConfigV2PolicyOrchestratorForOrganization#gcs}
	Gcs *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSourceGcs `field:"optional" json:"gcs" yaml:"gcs"`
	// A local path within the VM to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#local_path GoogleOsConfigV2PolicyOrchestratorForOrganization#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator_for_organization#remote GoogleOsConfigV2PolicyOrchestratorForOrganization#remote}
	Remote *GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmSourceRemote `field:"optional" json:"remote" yaml:"remote"`
}

