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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_topic#allowed_persistence_regions GooglePubsubTopic#allowed_persistence_regions}
	AllowedPersistenceRegions *[]*string `field:"required" json:"allowedPersistenceRegions" yaml:"allowedPersistenceRegions"`
}
