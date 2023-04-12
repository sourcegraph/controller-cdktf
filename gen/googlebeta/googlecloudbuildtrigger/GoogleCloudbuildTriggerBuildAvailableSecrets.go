package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildAvailableSecrets struct {
	// secret_manager block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#secret_manager GoogleCloudbuildTrigger#secret_manager}
	SecretManager interface{} `field:"required" json:"secretManager" yaml:"secretManager"`
}

