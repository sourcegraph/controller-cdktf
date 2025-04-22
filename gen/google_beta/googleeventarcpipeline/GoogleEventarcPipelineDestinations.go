package googleeventarcpipeline


type GoogleEventarcPipelineDestinations struct {
	// authentication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#authentication_config GoogleEventarcPipeline#authentication_config}
	AuthenticationConfig *GoogleEventarcPipelineDestinationsAuthenticationConfig `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// http_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#http_endpoint GoogleEventarcPipeline#http_endpoint}
	HttpEndpoint *GoogleEventarcPipelineDestinationsHttpEndpoint `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// The resource name of the Message Bus to which events should be published.
	//
	// The Message Bus resource should exist in the same project as
	// the Pipeline. Format:
	// 'projects/{project}/locations/{location}/messageBuses/{message_bus}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#message_bus GoogleEventarcPipeline#message_bus}
	MessageBus *string `field:"optional" json:"messageBus" yaml:"messageBus"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#network_config GoogleEventarcPipeline#network_config}
	NetworkConfig *GoogleEventarcPipelineDestinationsNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// output_payload_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#output_payload_format GoogleEventarcPipeline#output_payload_format}
	OutputPayloadFormat *GoogleEventarcPipelineDestinationsOutputPayloadFormat `field:"optional" json:"outputPayloadFormat" yaml:"outputPayloadFormat"`
	// The resource name of the Pub/Sub topic to which events should be published. Format: 'projects/{project}/locations/{location}/topics/{topic}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#topic GoogleEventarcPipeline#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same
	// project as the Pipeline. Format:
	// 'projects/{project}/locations/{location}/workflows/{workflow}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#workflow GoogleEventarcPipeline#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

