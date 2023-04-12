package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesPackageRepositoriesGoo struct {
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_guest_policies#name GoogleOsConfigGuestPolicies#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_guest_policies#url GoogleOsConfigGuestPolicies#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

