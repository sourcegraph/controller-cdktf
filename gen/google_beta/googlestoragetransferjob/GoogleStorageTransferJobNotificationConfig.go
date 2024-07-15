package googlestoragetransferjob


type GoogleStorageTransferJobNotificationConfig struct {
	// The desired format of the notification message payloads. One of "NONE" or "JSON".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#payload_format GoogleStorageTransferJob#payload_format}
	PayloadFormat *string `field:"required" json:"payloadFormat" yaml:"payloadFormat"`
	// The Topic.name of the Pub/Sub topic to which to publish notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#pubsub_topic GoogleStorageTransferJob#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// Event types for which a notification is desired.
	//
	// If empty, send notifications for all event types. The valid types are "TRANSFER_OPERATION_SUCCESS", "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#event_types GoogleStorageTransferJob#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
}

