package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec struct {
	// validate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#validate GoogleOsConfigOsPolicyAssignment#validate}
	Validate *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate `field:"required" json:"validate" yaml:"validate"`
	// enforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#enforce GoogleOsConfigOsPolicyAssignment#enforce}
	Enforce *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce `field:"optional" json:"enforce" yaml:"enforce"`
}

