package googlepubsubtopic


type GooglePubsubTopicMessageTransformsAiInferenceUnstructuredInference struct {
	// A parameters object to be included in each inference request.
	//
	// The parameters object is combined with the data field of the Pub/Sub
	// message to form the inference request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_topic#parameters GooglePubsubTopic#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

