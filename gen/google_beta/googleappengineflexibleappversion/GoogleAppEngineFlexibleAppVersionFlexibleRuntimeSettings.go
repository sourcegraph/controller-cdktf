package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionFlexibleRuntimeSettings struct {
	// Operating System of the application runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#operating_system GoogleAppEngineFlexibleAppVersion#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The runtime version of an App Engine flexible application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#runtime_version GoogleAppEngineFlexibleAppVersion#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

