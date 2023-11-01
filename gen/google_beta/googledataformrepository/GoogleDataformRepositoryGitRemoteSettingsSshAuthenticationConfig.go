package googledataformrepository


type GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig struct {
	// Content of a public SSH key to verify an identity of a remote Git host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_dataform_repository#host_public_key GoogleDataformRepository#host_public_key}
	HostPublicKey *string `field:"required" json:"hostPublicKey" yaml:"hostPublicKey"`
	// The name of the Secret Manager secret version to use as a ssh private key for Git operations.
	//
	// Must be in the format projects/*\/secrets/*\/versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_dataform_repository#user_private_key_secret_version GoogleDataformRepository#user_private_key_secret_version}
	UserPrivateKeySecretVersion *string `field:"required" json:"userPrivateKeySecretVersion" yaml:"userPrivateKeySecretVersion"`
}

