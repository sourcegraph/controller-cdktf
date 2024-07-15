package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesPackageRepositoriesApt struct {
	// List of components for this repository. Must contain at least one item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#components GoogleOsConfigGuestPolicies#components}
	Components *[]*string `field:"required" json:"components" yaml:"components"`
	// Distribution of this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#distribution GoogleOsConfigGuestPolicies#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// URI for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#uri GoogleOsConfigGuestPolicies#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Type of archive files in this repository. The default behavior is DEB. Default value: "DEB" Possible values: ["DEB", "DEB_SRC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#archive_type GoogleOsConfigGuestPolicies#archive_type}
	ArchiveType *string `field:"optional" json:"archiveType" yaml:"archiveType"`
	// URI of the key file for this repository.
	//
	// The agent maintains a keyring at
	// /etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg containing all the keys in any applied guest policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#gpg_key GoogleOsConfigGuestPolicies#gpg_key}
	GpgKey *string `field:"optional" json:"gpgKey" yaml:"gpgKey"`
}

