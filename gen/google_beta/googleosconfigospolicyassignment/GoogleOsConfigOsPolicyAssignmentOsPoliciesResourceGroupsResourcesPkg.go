package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg struct {
	// The desired state the agent should maintain for this package. Possible values: ["DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#desired_state GoogleOsConfigOsPolicyAssignment#desired_state}
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#apt GoogleOsConfigOsPolicyAssignment#apt}
	Apt *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt `field:"optional" json:"apt" yaml:"apt"`
	// deb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#deb GoogleOsConfigOsPolicyAssignment#deb}
	Deb *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb `field:"optional" json:"deb" yaml:"deb"`
	// googet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#googet GoogleOsConfigOsPolicyAssignment#googet}
	Googet *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget `field:"optional" json:"googet" yaml:"googet"`
	// msi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#msi GoogleOsConfigOsPolicyAssignment#msi}
	Msi *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi `field:"optional" json:"msi" yaml:"msi"`
	// rpm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#rpm GoogleOsConfigOsPolicyAssignment#rpm}
	Rpm *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm `field:"optional" json:"rpm" yaml:"rpm"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#yum GoogleOsConfigOsPolicyAssignment#yum}
	Yum *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#zypper GoogleOsConfigOsPolicyAssignment#zypper}
	Zypper *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

