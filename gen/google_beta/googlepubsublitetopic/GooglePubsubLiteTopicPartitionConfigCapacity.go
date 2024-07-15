package googlepubsublitetopic


type GooglePubsubLiteTopicPartitionConfigCapacity struct {
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_lite_topic#publish_mib_per_sec GooglePubsubLiteTopic#publish_mib_per_sec}
	PublishMibPerSec *float64 `field:"required" json:"publishMibPerSec" yaml:"publishMibPerSec"`
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_lite_topic#subscribe_mib_per_sec GooglePubsubLiteTopic#subscribe_mib_per_sec}
	SubscribeMibPerSec *float64 `field:"required" json:"subscribeMibPerSec" yaml:"subscribeMibPerSec"`
}

