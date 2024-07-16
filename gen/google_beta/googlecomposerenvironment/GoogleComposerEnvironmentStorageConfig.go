package googlecomposerenvironment


type GoogleComposerEnvironmentStorageConfig struct {
	// Optional. Name of an existing Cloud Storage bucket to be used by the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#bucket GoogleComposerEnvironment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
}

