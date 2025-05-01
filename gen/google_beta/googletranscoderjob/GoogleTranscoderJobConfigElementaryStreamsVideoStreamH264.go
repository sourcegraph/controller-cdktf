package googletranscoderjob


type GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264 struct {
	// The video bitrate in bits per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#bitrate_bps GoogleTranscoderJob#bitrate_bps}
	BitrateBps *float64 `field:"required" json:"bitrateBps" yaml:"bitrateBps"`
	// The target video frame rate in frames per second (FPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#frame_rate GoogleTranscoderJob#frame_rate}
	FrameRate *float64 `field:"required" json:"frameRate" yaml:"frameRate"`
	// Target CRF level. The default is '21'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#crf_level GoogleTranscoderJob#crf_level}
	CrfLevel *float64 `field:"optional" json:"crfLevel" yaml:"crfLevel"`
	// The entropy coder to use. The default is 'cabac'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#entropy_coder GoogleTranscoderJob#entropy_coder}
	EntropyCoder *string `field:"optional" json:"entropyCoder" yaml:"entropyCoder"`
	// Select the GOP size based on the specified duration. The default is '3s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#gop_duration GoogleTranscoderJob#gop_duration}
	GopDuration *string `field:"optional" json:"gopDuration" yaml:"gopDuration"`
	// The height of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#height_pixels GoogleTranscoderJob#height_pixels}
	HeightPixels *float64 `field:"optional" json:"heightPixels" yaml:"heightPixels"`
	// hlg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#hlg GoogleTranscoderJob#hlg}
	Hlg *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Hlg `field:"optional" json:"hlg" yaml:"hlg"`
	// Pixel format to use. The default is 'yuv420p'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#pixel_format GoogleTranscoderJob#pixel_format}
	PixelFormat *string `field:"optional" json:"pixelFormat" yaml:"pixelFormat"`
	// Enforces the specified codec preset. The default is 'veryfast'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#preset GoogleTranscoderJob#preset}
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Enforces the specified codec profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#profile GoogleTranscoderJob#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Specify the mode. The default is 'vbr'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#rate_control_mode GoogleTranscoderJob#rate_control_mode}
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// sdr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#sdr GoogleTranscoderJob#sdr}
	Sdr *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264Sdr `field:"optional" json:"sdr" yaml:"sdr"`
	// Initial fullness of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#vbv_fullness_bits GoogleTranscoderJob#vbv_fullness_bits}
	VbvFullnessBits *float64 `field:"optional" json:"vbvFullnessBits" yaml:"vbvFullnessBits"`
	// Size of the Video Buffering Verifier (VBV) buffer in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#vbv_size_bits GoogleTranscoderJob#vbv_size_bits}
	VbvSizeBits *float64 `field:"optional" json:"vbvSizeBits" yaml:"vbvSizeBits"`
	// The width of the video in pixels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#width_pixels GoogleTranscoderJob#width_pixels}
	WidthPixels *float64 `field:"optional" json:"widthPixels" yaml:"widthPixels"`
}

