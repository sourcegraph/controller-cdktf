package googletranscoderjob


type GoogleTranscoderJobConfigMuxStreams struct {
	// The container format. The default is 'mp4'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#container GoogleTranscoderJob#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// List of ElementaryStream.key values multiplexed in this stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#elementary_streams GoogleTranscoderJob#elementary_streams}
	ElementaryStreams *[]*string `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// Identifier of the encryption configuration to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#encryption_id GoogleTranscoderJob#encryption_id}
	EncryptionId *string `field:"optional" json:"encryptionId" yaml:"encryptionId"`
	// The name of the generated file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#file_name GoogleTranscoderJob#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// A unique key for this multiplexed stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#key GoogleTranscoderJob#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// segment_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#segment_settings GoogleTranscoderJob#segment_settings}
	SegmentSettings *GoogleTranscoderJobConfigMuxStreamsSegmentSettings `field:"optional" json:"segmentSettings" yaml:"segmentSettings"`
}

