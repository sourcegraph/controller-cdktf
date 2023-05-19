package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsRtmpOutputSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsRtmpOutputSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#certficate_mode MedialiveChannel#certficate_mode}.
	CertficateMode *string `field:"optional" json:"certficateMode" yaml:"certficateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#connection_retry_interval MedialiveChannel#connection_retry_interval}.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#num_retries MedialiveChannel#num_retries}.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}

