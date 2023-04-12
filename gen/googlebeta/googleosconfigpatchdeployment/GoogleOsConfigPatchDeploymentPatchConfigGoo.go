package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigGoo struct {
	// goo update settings. Use this setting to override the default goo patch rules.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#enabled GoogleOsConfigPatchDeployment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

