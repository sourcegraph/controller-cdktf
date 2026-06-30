package kinesisfirehosedeliverystream

type KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kinesis_firehose_delivery_stream#enabled KinesisFirehoseDeliveryStream#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// processors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kinesis_firehose_delivery_stream#processors KinesisFirehoseDeliveryStream#processors}
	Processors any `field:"optional" json:"processors" yaml:"processors"`
}
