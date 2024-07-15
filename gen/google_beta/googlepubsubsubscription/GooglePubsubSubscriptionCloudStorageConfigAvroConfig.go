package googlepubsubsubscription


type GooglePubsubSubscriptionCloudStorageConfigAvroConfig struct {
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey as additional fields in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription#write_metadata GooglePubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

