package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigMuxStreams struct {
	// The container format. The default is 'mp4'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#container GoogleTranscoderJobTemplate#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// List of ElementaryStream.key values multiplexed in this stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#elementary_streams GoogleTranscoderJobTemplate#elementary_streams}
	ElementaryStreams *[]*string `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// Identifier of the encryption configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#encryption_id GoogleTranscoderJobTemplate#encryption_id}
	EncryptionId *string `field:"optional" json:"encryptionId" yaml:"encryptionId"`
	// The name of the generated file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#file_name GoogleTranscoderJobTemplate#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// A unique key for this multiplexed stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#key GoogleTranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// segment_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#segment_settings GoogleTranscoderJobTemplate#segment_settings}
	SegmentSettings *GoogleTranscoderJobTemplateConfigMuxStreamsSegmentSettings `field:"optional" json:"segmentSettings" yaml:"segmentSettings"`
}

