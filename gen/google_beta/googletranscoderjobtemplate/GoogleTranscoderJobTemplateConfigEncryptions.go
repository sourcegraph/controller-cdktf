package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigEncryptions struct {
	// Identifier for this set of encryption options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#id GoogleTranscoderJobTemplate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// aes128 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#aes128 GoogleTranscoderJobTemplate#aes128}
	Aes128 *GoogleTranscoderJobTemplateConfigEncryptionsAes128 `field:"optional" json:"aes128" yaml:"aes128"`
	// drm_systems block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#drm_systems GoogleTranscoderJobTemplate#drm_systems}
	DrmSystems *GoogleTranscoderJobTemplateConfigEncryptionsDrmSystems `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// mpeg_cenc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#mpeg_cenc GoogleTranscoderJobTemplate#mpeg_cenc}
	MpegCenc *GoogleTranscoderJobTemplateConfigEncryptionsMpegCenc `field:"optional" json:"mpegCenc" yaml:"mpegCenc"`
	// sample_aes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#sample_aes GoogleTranscoderJobTemplate#sample_aes}
	SampleAes *GoogleTranscoderJobTemplateConfigEncryptionsSampleAes `field:"optional" json:"sampleAes" yaml:"sampleAes"`
	// secret_manager_key_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#secret_manager_key_source GoogleTranscoderJobTemplate#secret_manager_key_source}
	SecretManagerKeySource *GoogleTranscoderJobTemplateConfigEncryptionsSecretManagerKeySource `field:"optional" json:"secretManagerKeySource" yaml:"secretManagerKeySource"`
}

