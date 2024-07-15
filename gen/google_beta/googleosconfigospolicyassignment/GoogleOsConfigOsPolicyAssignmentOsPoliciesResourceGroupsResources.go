package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResources struct {
	// The id of the resource with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens.
	//
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the OS policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#id GoogleOsConfigOsPolicyAssignment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#exec GoogleOsConfigOsPolicyAssignment#exec}
	Exec *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec `field:"optional" json:"exec" yaml:"exec"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#file GoogleOsConfigOsPolicyAssignment#file}
	File *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile `field:"optional" json:"file" yaml:"file"`
	// pkg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#pkg GoogleOsConfigOsPolicyAssignment#pkg}
	Pkg *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg `field:"optional" json:"pkg" yaml:"pkg"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#repository GoogleOsConfigOsPolicyAssignment#repository}
	Repository *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository `field:"optional" json:"repository" yaml:"repository"`
}

