package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems struct {
	// clearkey block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#clearkey GoogleTranscoderJobTemplate#clearkey}
	Clearkey *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey `field:"optional" json:"clearkey" yaml:"clearkey"`
	// fairplay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#fairplay GoogleTranscoderJobTemplate#fairplay}
	Fairplay *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay `field:"optional" json:"fairplay" yaml:"fairplay"`
	// playready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#playready GoogleTranscoderJobTemplate#playready}
	Playready *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready `field:"optional" json:"playready" yaml:"playready"`
	// widevine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#widevine GoogleTranscoderJobTemplate#widevine}
	Widevine *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine `field:"optional" json:"widevine" yaml:"widevine"`
}

