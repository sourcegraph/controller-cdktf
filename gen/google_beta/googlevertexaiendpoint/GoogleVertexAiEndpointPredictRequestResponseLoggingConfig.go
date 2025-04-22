package googlevertexaiendpoint


type GoogleVertexAiEndpointPredictRequestResponseLoggingConfig struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_endpoint#bigquery_destination GoogleVertexAiEndpoint#bigquery_destination}
	BigqueryDestination *GoogleVertexAiEndpointPredictRequestResponseLoggingConfigBigqueryDestination `field:"optional" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// If logging is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_endpoint#enabled GoogleVertexAiEndpoint#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Percentage of requests to be logged, expressed as a fraction in range(0,1].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_endpoint#sampling_rate GoogleVertexAiEndpoint#sampling_rate}
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

