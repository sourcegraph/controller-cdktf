package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#source GoogleOsConfigOsPolicyAssignment#source}
	Source *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSource `field:"required" json:"source" yaml:"source"`
	// Additional properties to use during installation.
	//
	// This should be in the format of Property=Setting. Appended to the defaults of 'ACTION=INSTALL REBOOT=ReallySuppress'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#properties GoogleOsConfigOsPolicyAssignment#properties}
	Properties *[]*string `field:"optional" json:"properties" yaml:"properties"`
}

