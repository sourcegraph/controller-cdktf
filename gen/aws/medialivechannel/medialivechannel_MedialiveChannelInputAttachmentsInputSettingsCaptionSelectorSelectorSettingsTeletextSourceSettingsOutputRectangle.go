package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsOutputRectangle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#height MedialiveChannel#height}.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#left_offset MedialiveChannel#left_offset}.
	LeftOffset *float64 `field:"required" json:"leftOffset" yaml:"leftOffset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#top_offset MedialiveChannel#top_offset}.
	TopOffset *float64 `field:"required" json:"topOffset" yaml:"topOffset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#width MedialiveChannel#width}.
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

