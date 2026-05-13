package pubsubsubscription


type PubsubSubscriptionMessageTransformsAiInferenceUnstructuredInference struct {
	// A parameters object to be included in each inference request.
	//
	// The parameters object is combined with the data field of the Pub/Sub
	// message to form the inference request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#parameters PubsubSubscription#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

