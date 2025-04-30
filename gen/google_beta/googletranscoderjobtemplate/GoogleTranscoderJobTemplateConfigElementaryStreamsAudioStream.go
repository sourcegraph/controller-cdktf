package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream struct {
	// Audio bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#bitrate_bps GoogleTranscoderJobTemplate#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// Number of audio channels. The default is '2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#channel_count GoogleTranscoderJobTemplate#channel_count}
	ChannelCount *float64 `field:"optional" json:"channelCount" yaml:"channelCount"`
	// A list of channel names specifying layout of the audio channels.  The default is ["fl", "fr"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#channel_layout GoogleTranscoderJobTemplate#channel_layout}
	ChannelLayout *[]*string `field:"optional" json:"channelLayout" yaml:"channelLayout"`
	// The codec for this audio stream. The default is 'aac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#codec GoogleTranscoderJobTemplate#codec}
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The audio sample rate in Hertz. The default is '48000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#sample_rate_hertz GoogleTranscoderJobTemplate#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
}

