package googlecomposerenvironment


type GoogleComposerEnvironmentConfigEncryptionConfig struct {
	// Optional. Customer-managed Encryption Key available through Google's Key Management Service. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#kms_key_name GoogleComposerEnvironment#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

