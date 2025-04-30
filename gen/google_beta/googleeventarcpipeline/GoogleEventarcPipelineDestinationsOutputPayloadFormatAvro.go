package googleeventarcpipeline


type GoogleEventarcPipelineDestinationsOutputPayloadFormatAvro struct {
	// The entire schema definition is stored in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#schema_definition GoogleEventarcPipeline#schema_definition}
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

