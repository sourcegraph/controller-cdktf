package googleeventarcpipeline


type GoogleEventarcPipelineDestinationsOutputPayloadFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#avro GoogleEventarcPipeline#avro}
	Avro *GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#json GoogleEventarcPipeline#json}
	Json *GoogleEventarcPipelineDestinationsOutputPayloadFormatJson `field:"optional" json:"json" yaml:"json"`
	// protobuf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#protobuf GoogleEventarcPipeline#protobuf}
	Protobuf *GoogleEventarcPipelineDestinationsOutputPayloadFormatProtobuf `field:"optional" json:"protobuf" yaml:"protobuf"`
}

