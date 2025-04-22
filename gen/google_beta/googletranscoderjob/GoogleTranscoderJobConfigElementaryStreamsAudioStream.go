package googletranscoderjob


type GoogleTranscoderJobConfigElementaryStreamsAudioStream struct {
	// Audio bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#bitrate_bps GoogleTranscoderJob#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// Number of audio channels. The default is '2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#channel_count GoogleTranscoderJob#channel_count}
	ChannelCount *float64 `field:"optional" json:"channelCount" yaml:"channelCount"`
	// A list of channel names specifying layout of the audio channels. The default is ["fl", "fr"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#channel_layout GoogleTranscoderJob#channel_layout}
	ChannelLayout *[]*string `field:"optional" json:"channelLayout" yaml:"channelLayout"`
	// The codec for this audio stream. The default is 'aac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#codec GoogleTranscoderJob#codec}
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// The audio sample rate in Hertz. The default is '48000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#sample_rate_hertz GoogleTranscoderJob#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
}

