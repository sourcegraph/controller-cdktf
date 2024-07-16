package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildAvailableSecretsSecretManager struct {
	// Environment variable name to associate with the secret.
	//
	// Secret environment
	// variables must be unique across all of a build's secrets, and must be used
	// by at least one build step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#env GoogleCloudbuildTrigger#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// Resource name of the SecretVersion. In format: projects/* /secrets/* /versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#version_name GoogleCloudbuildTrigger#version_name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
}

