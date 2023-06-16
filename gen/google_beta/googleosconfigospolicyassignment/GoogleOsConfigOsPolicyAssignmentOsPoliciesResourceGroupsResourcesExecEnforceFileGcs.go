package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileGcs struct {
	// Required. Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_os_config_os_policy_assignment#bucket GoogleOsConfigOsPolicyAssignment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Required. Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_os_config_os_policy_assignment#object GoogleOsConfigOsPolicyAssignment#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_os_config_os_policy_assignment#generation GoogleOsConfigOsPolicyAssignment#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
}

