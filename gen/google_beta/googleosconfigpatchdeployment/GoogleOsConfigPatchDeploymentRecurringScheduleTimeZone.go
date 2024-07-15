package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentRecurringScheduleTimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#id GoogleOsConfigPatchDeployment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// IANA Time Zone Database version number, e.g. "2019a".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#version GoogleOsConfigPatchDeployment#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

