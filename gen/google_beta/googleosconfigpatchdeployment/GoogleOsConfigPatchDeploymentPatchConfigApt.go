package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigApt struct {
	// List of packages to exclude from update. These packages will be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#excludes GoogleOsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of packages to be updated.
	//
	// These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#exclusive_packages GoogleOsConfigPatchDeployment#exclusive_packages}
	ExclusivePackages *[]*string `field:"optional" json:"exclusivePackages" yaml:"exclusivePackages"`
	// By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: ["DIST", "UPGRADE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#type GoogleOsConfigPatchDeployment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

