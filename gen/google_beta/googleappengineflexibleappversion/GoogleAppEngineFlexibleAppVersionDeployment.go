package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionDeployment struct {
	// cloud_build_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#cloud_build_options GoogleAppEngineFlexibleAppVersion#cloud_build_options}
	CloudBuildOptions *GoogleAppEngineFlexibleAppVersionDeploymentCloudBuildOptions `field:"optional" json:"cloudBuildOptions" yaml:"cloudBuildOptions"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#container GoogleAppEngineFlexibleAppVersion#container}
	Container *GoogleAppEngineFlexibleAppVersionDeploymentContainer `field:"optional" json:"container" yaml:"container"`
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#files GoogleAppEngineFlexibleAppVersion#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// zip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_flexible_app_version#zip GoogleAppEngineFlexibleAppVersion#zip}
	Zip *GoogleAppEngineFlexibleAppVersionDeploymentZip `field:"optional" json:"zip" yaml:"zip"`
}

