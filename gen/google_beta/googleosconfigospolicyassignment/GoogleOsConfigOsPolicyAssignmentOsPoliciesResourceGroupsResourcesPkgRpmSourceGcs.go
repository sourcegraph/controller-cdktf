package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSourceGcs struct {
	// Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#bucket GoogleOsConfigOsPolicyAssignment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#object GoogleOsConfigOsPolicyAssignment#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#generation GoogleOsConfigOsPolicyAssignment#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
}

