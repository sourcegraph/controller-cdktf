package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigElementaryStreams struct {
	// audio_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#audio_stream GoogleTranscoderJobTemplate#audio_stream}
	AudioStream *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream `field:"optional" json:"audioStream" yaml:"audioStream"`
	// A unique key for this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#key GoogleTranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// video_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#video_stream GoogleTranscoderJobTemplate#video_stream}
	VideoStream *GoogleTranscoderJobTemplateConfigElementaryStreamsVideoStream `field:"optional" json:"videoStream" yaml:"videoStream"`
}

