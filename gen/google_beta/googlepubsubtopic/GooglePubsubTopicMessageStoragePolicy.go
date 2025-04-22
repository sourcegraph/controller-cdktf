package googlepubsubtopic


type GooglePubsubTopicMessageStoragePolicy struct {
	// A list of IDs of GCP regions where messages that are published to the topic may be persisted in storage.
	//
	// Messages published by
	// publishers running in non-allowed GCP regions (or running outside
	// of GCP altogether) will be routed for storage in one of the
	// allowed regions. An empty list means that no regions are allowed,
	// and is not a valid configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#allowed_persistence_regions GooglePubsubTopic#allowed_persistence_regions}
	AllowedPersistenceRegions *[]*string `field:"required" json:"allowedPersistenceRegions" yaml:"allowedPersistenceRegions"`
	// If true, 'allowedPersistenceRegions' is also used to enforce in-transit guarantees for messages.
	//
	// That is, Pub/Sub will fail topics.publish
	// operations on this topic and subscribe operations on any subscription
	// attached to this topic in any region that is not in 'allowedPersistenceRegions'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_pubsub_topic#enforce_in_transit GooglePubsubTopic#enforce_in_transit}
	EnforceInTransit interface{} `field:"optional" json:"enforceInTransit" yaml:"enforceInTransit"`
}

