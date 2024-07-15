package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesPackages struct {
	// The name of the package.
	//
	// A package is uniquely identified for conflict validation
	// by checking the package name and the manager(s) that the package targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#name GoogleOsConfigGuestPolicies#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The desiredState the agent should maintain for this package.
	//
	// The default is to ensure the package is installed. Possible values: ["INSTALLED", "UPDATED", "REMOVED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#desired_state GoogleOsConfigGuestPolicies#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Type of package manager that can be used to install this package.
	//
	// If a system does not have the package manager,
	// the package is not installed or removed no error message is returned. By default, or if you specify ANY,
	// the agent attempts to install and remove this package using the default package manager.
	// This is useful when creating a policy that applies to different types of systems.
	// The default behavior is ANY. Default value: "ANY" Possible values: ["ANY", "APT", "YUM", "ZYPPER", "GOO"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#manager GoogleOsConfigGuestPolicies#manager}
	Manager *string `field:"optional" json:"manager" yaml:"manager"`
}

