package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionEntrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#shell GoogleAppEngineStandardAppVersion#shell}
	Shell *string `field:"required" json:"shell" yaml:"shell"`
}

