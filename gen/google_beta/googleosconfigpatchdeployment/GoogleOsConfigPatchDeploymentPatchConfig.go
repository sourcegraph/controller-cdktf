package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfig struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#apt GoogleOsConfigPatchDeployment#apt}
	Apt *GoogleOsConfigPatchDeploymentPatchConfigApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#goo GoogleOsConfigPatchDeployment#goo}
	Goo *GoogleOsConfigPatchDeploymentPatchConfigGoo `field:"optional" json:"goo" yaml:"goo"`
	// Allows the patch job to run on Managed instance groups (MIGs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#mig_instances_allowed GoogleOsConfigPatchDeployment#mig_instances_allowed}
	MigInstancesAllowed interface{} `field:"optional" json:"migInstancesAllowed" yaml:"migInstancesAllowed"`
	// post_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#post_step GoogleOsConfigPatchDeployment#post_step}
	PostStep *GoogleOsConfigPatchDeploymentPatchConfigPostStep `field:"optional" json:"postStep" yaml:"postStep"`
	// pre_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#pre_step GoogleOsConfigPatchDeployment#pre_step}
	PreStep *GoogleOsConfigPatchDeploymentPatchConfigPreStep `field:"optional" json:"preStep" yaml:"preStep"`
	// Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#reboot_config GoogleOsConfigPatchDeployment#reboot_config}
	RebootConfig *string `field:"optional" json:"rebootConfig" yaml:"rebootConfig"`
	// windows_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#windows_update GoogleOsConfigPatchDeployment#windows_update}
	WindowsUpdate *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate `field:"optional" json:"windowsUpdate" yaml:"windowsUpdate"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#yum GoogleOsConfigPatchDeployment#yum}
	Yum *GoogleOsConfigPatchDeploymentPatchConfigYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#zypper GoogleOsConfigPatchDeployment#zypper}
	Zypper *GoogleOsConfigPatchDeploymentPatchConfigZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

