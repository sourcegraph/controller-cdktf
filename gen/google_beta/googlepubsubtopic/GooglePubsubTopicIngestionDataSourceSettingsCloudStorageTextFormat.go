package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat struct {
	// The delimiter to use when using the 'text' format.
	//
	// Each line of text as
	// specified by the delimiter will be set to the 'data' field of a Pub/Sub
	// message. When unset, '\n' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#delimiter GooglePubsubTopic#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

