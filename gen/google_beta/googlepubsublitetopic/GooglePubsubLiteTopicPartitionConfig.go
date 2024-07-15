package googlepubsublitetopic


type GooglePubsubLiteTopicPartitionConfig struct {
	// The number of partitions in the topic. Must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_lite_topic#count GooglePubsubLiteTopic#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_lite_topic#capacity GooglePubsubLiteTopic#capacity}
	Capacity *GooglePubsubLiteTopicPartitionConfigCapacity `field:"optional" json:"capacity" yaml:"capacity"`
}

