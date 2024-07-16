package googlepubsubsubscription


type GooglePubsubSubscriptionPushConfigNoWrapper struct {
	// When true, writes the Pub/Sub message metadata to 'x-goog-pubsub-<KEY>:<VAL>' headers of the HTTP request.
	//
	// Writes the
	// Pub/Sub message attributes to '<KEY>:<VAL>' headers of the HTTP request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription#write_metadata GooglePubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"required" json:"writeMetadata" yaml:"writeMetadata"`
}

