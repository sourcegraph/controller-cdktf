package googlepubsubtopic


type GooglePubsubTopicMessageTransformsAiInference struct {
	// The endpoint to a Vertex AI model of the form 'projects/{project}/locations/{location}/endpoints/{endpoint}' or 'projects/{project}/locations/{location}/publishers/{publisher}/models/{model}'.
	//
	// Vertex AI API requests will be sent to this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_topic#endpoint GooglePubsubTopic#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The service account to use to make prediction requests against endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_topic#service_account_email GooglePubsubTopic#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// unstructured_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_topic#unstructured_inference GooglePubsubTopic#unstructured_inference}
	UnstructuredInference *GooglePubsubTopicMessageTransformsAiInferenceUnstructuredInference `field:"optional" json:"unstructuredInference" yaml:"unstructuredInference"`
}

