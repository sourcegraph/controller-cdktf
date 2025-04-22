package googletranscoderjob


type GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource struct {
	// The name of the Secret Version containing the encryption key in the following format: projects/{project}/secrets/{secret_id}/versions/{version_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#secret_version GoogleTranscoderJob#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

