package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsRemixSettings struct {
	// channel_mappings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#channel_mappings MedialiveChannel#channel_mappings}
	ChannelMappings interface{} `field:"required" json:"channelMappings" yaml:"channelMappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#channels_in MedialiveChannel#channels_in}.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#channels_out MedialiveChannel#channels_out}.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

