package cloudfrontrealtimelogconfig


type CloudfrontRealtimeLogConfigEndpoint struct {
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config#kinesis_stream_config CloudfrontRealtimeLogConfig#kinesis_stream_config}
	KinesisStreamConfig *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config#stream_type CloudfrontRealtimeLogConfig#stream_type}.
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}

