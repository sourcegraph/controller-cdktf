package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentOneTimeSchedule struct {
	// The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#execute_time GoogleOsConfigPatchDeployment#execute_time}
	ExecuteTime *string `field:"required" json:"executeTime" yaml:"executeTime"`
}
