package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionEntrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_flexible_app_version#shell GoogleAppEngineFlexibleAppVersion#shell}
	Shell *string `field:"required" json:"shell" yaml:"shell"`
}

