package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile struct {
	// The absolute path of the file within the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#path GoogleOsConfigOsPolicyAssignment#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Desired state of the file. Possible values: ["DESIRED_STATE_UNSPECIFIED", "PRESENT", "ABSENT", "CONTENTS_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#state GoogleOsConfigOsPolicyAssignment#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// A a file with this content. The size of the content is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#content GoogleOsConfigOsPolicyAssignment#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#file GoogleOsConfigOsPolicyAssignment#file}
	File *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileFile `field:"optional" json:"file" yaml:"file"`
}

