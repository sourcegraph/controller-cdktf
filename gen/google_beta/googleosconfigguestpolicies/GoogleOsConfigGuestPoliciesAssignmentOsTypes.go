package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesAssignmentOsTypes struct {
	// Targets VM instances with OS Inventory enabled and having the following OS architecture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#os_architecture GoogleOsConfigGuestPolicies#os_architecture}
	OsArchitecture *string `field:"optional" json:"osArchitecture" yaml:"osArchitecture"`
	// Targets VM instances with OS Inventory enabled and having the following OS short name, for example "debian" or "windows".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#os_short_name GoogleOsConfigGuestPolicies#os_short_name}
	OsShortName *string `field:"optional" json:"osShortName" yaml:"osShortName"`
	// Targets VM instances with OS Inventory enabled and having the following following OS version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#os_version GoogleOsConfigGuestPolicies#os_version}
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
}

