package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentPatchConfigGoo struct {
	// goo update settings. Use this setting to override the default goo patch rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#enabled GoogleOsConfigPatchDeployment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

