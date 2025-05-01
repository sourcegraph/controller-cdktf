package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource struct {
	// The name of the Secret Version containing the encryption key in the following format: projects/{project}/secrets/{secret_id}/versions/{version_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#secret_version GoogleTranscoderJobTemplate#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

