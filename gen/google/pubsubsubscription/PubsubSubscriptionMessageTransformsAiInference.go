package pubsubsubscription


type PubsubSubscriptionMessageTransformsAiInference struct {
	// The endpoint to a Vertex AI model of the form 'projects/{project}/locations/{location}/endpoints/{endpoint}' or 'projects/{project}/locations/{location}/publishers/{publisher}/models/{model}'.
	//
	// Vertex AI API requests will be sent to this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#endpoint PubsubSubscription#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The service account to use to make prediction requests against endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#service_account_email PubsubSubscription#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// unstructured_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#unstructured_inference PubsubSubscription#unstructured_inference}
	UnstructuredInference *PubsubSubscriptionMessageTransformsAiInferenceUnstructuredInference `field:"optional" json:"unstructuredInference" yaml:"unstructuredInference"`
}

