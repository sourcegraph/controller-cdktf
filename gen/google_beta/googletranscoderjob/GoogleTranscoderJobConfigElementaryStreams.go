package googletranscoderjob


type GoogleTranscoderJobConfigElementaryStreams struct {
	// audio_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#audio_stream GoogleTranscoderJob#audio_stream}
	AudioStream *GoogleTranscoderJobConfigElementaryStreamsAudioStream `field:"optional" json:"audioStream" yaml:"audioStream"`
	// A unique key for this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#key GoogleTranscoderJob#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// video_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#video_stream GoogleTranscoderJob#video_stream}
	VideoStream *GoogleTranscoderJobConfigElementaryStreamsVideoStream `field:"optional" json:"videoStream" yaml:"videoStream"`
}

