package pubsubtopic


type PubsubTopicMessageTransformsAiInference struct {
	// The endpoint to a Vertex AI model of the form 'projects/{project}/locations/{location}/endpoints/{endpoint}' or 'projects/{project}/locations/{location}/publishers/{publisher}/models/{model}'.
	//
	// Vertex AI API requests will be sent to this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#endpoint PubsubTopic#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The service account to use to make prediction requests against endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#service_account_email PubsubTopic#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// unstructured_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#unstructured_inference PubsubTopic#unstructured_inference}
	UnstructuredInference *PubsubTopicMessageTransformsAiInferenceUnstructuredInference `field:"optional" json:"unstructuredInference" yaml:"unstructuredInference"`
}

