package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions struct {
	// Path to the yaml file used in deployment, used to determine runtime configuration details.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_flexible_app_version#app_yaml_path GoogleAppEngineFlexibleAppVersion#app_yaml_path}
	AppYamlPath *string `field:"required" json:"appYamlPath" yaml:"appYamlPath"`
	// The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_flexible_app_version#cloud_build_timeout GoogleAppEngineFlexibleAppVersion#cloud_build_timeout}
	CloudBuildTimeout *string `field:"optional" json:"cloudBuildTimeout" yaml:"cloudBuildTimeout"`
}
