package googletranscoderjob


type GoogleTranscoderJobConfigEncryptions struct {
	// Identifier for this set of encryption options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#id GoogleTranscoderJob#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// aes128 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#aes128 GoogleTranscoderJob#aes128}
	Aes128 *GoogleTranscoderJobConfigEncryptionsAes128 `field:"optional" json:"aes128" yaml:"aes128"`
	// drm_systems block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#drm_systems GoogleTranscoderJob#drm_systems}
	DrmSystems *GoogleTranscoderJobConfigEncryptionsDrmSystems `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// mpeg_cenc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#mpeg_cenc GoogleTranscoderJob#mpeg_cenc}
	MpegCenc *GoogleTranscoderJobConfigEncryptionsMpegCenc `field:"optional" json:"mpegCenc" yaml:"mpegCenc"`
	// sample_aes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#sample_aes GoogleTranscoderJob#sample_aes}
	SampleAes *GoogleTranscoderJobConfigEncryptionsSampleAes `field:"optional" json:"sampleAes" yaml:"sampleAes"`
	// secret_manager_key_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#secret_manager_key_source GoogleTranscoderJob#secret_manager_key_source}
	SecretManagerKeySource *GoogleTranscoderJobConfigEncryptionsSecretManagerKeySource `field:"optional" json:"secretManagerKeySource" yaml:"secretManagerKeySource"`
}

