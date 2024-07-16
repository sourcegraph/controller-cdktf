package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesPackageRepositories struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#apt GoogleOsConfigGuestPolicies#apt}
	Apt *GoogleOsConfigGuestPoliciesPackageRepositoriesApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#goo GoogleOsConfigGuestPolicies#goo}
	Goo *GoogleOsConfigGuestPoliciesPackageRepositoriesGoo `field:"optional" json:"goo" yaml:"goo"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#yum GoogleOsConfigGuestPolicies#yum}
	Yum *GoogleOsConfigGuestPoliciesPackageRepositoriesYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#zypper GoogleOsConfigGuestPolicies#zypper}
	Zypper *GoogleOsConfigGuestPoliciesPackageRepositoriesZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

