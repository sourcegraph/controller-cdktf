package googletranscoderjob


type GoogleTranscoderJobConfigEncryptionsDrmSystems struct {
	// clearkey block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#clearkey GoogleTranscoderJob#clearkey}
	Clearkey *GoogleTranscoderJobConfigEncryptionsDrmSystemsClearkey `field:"optional" json:"clearkey" yaml:"clearkey"`
	// fairplay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#fairplay GoogleTranscoderJob#fairplay}
	Fairplay *GoogleTranscoderJobConfigEncryptionsDrmSystemsFairplay `field:"optional" json:"fairplay" yaml:"fairplay"`
	// playready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#playready GoogleTranscoderJob#playready}
	Playready *GoogleTranscoderJobConfigEncryptionsDrmSystemsPlayready `field:"optional" json:"playready" yaml:"playready"`
	// widevine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#widevine GoogleTranscoderJob#widevine}
	Widevine *GoogleTranscoderJobConfigEncryptionsDrmSystemsWidevine `field:"optional" json:"widevine" yaml:"widevine"`
}

