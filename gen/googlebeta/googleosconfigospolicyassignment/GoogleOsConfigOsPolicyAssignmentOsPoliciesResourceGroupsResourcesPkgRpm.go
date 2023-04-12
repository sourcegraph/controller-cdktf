package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#source GoogleOsConfigOsPolicyAssignment#source}
	Source *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: `rpm --upgrade --replacepkgs package.rpm` - install when true: `yum -y install package.rpm` or `zypper -y install package.rpm`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#pull_deps GoogleOsConfigOsPolicyAssignment#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

