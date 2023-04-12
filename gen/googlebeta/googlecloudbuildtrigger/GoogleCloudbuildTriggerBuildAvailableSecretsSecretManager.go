package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildAvailableSecretsSecretManager struct {
	// Environment variable name to associate with the secret.
	//
	// Secret environment
	// variables must be unique across all of a build's secrets, and must be used
	// by at least one build step.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#env GoogleCloudbuildTrigger#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// Resource name of the SecretVersion. In format: projects/*\/secrets/*\/versions/*.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#version_name GoogleCloudbuildTrigger#version_name}
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
}

