package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification struct {
	// How much data to include in the pub/sub message. Possible values: ["TABLE_PROFILE", "RESOURCE_NAME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#detail_of_message GoogleDataLossPreventionDiscoveryConfig#detail_of_message}
	DetailOfMessage *string `field:"optional" json:"detailOfMessage" yaml:"detailOfMessage"`
	// The type of event that triggers a Pub/Sub.
	//
	// At most one PubSubNotification per EventType is permitted. Possible values: ["NEW_PROFILE", "CHANGED_PROFILE", "SCORE_INCREASED", "ERROR_CHANGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#event GoogleDataLossPreventionDiscoveryConfig#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// pubsub_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#pubsub_condition GoogleDataLossPreventionDiscoveryConfig#pubsub_condition}
	PubsubCondition *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubCondition `field:"optional" json:"pubsubCondition" yaml:"pubsubCondition"`
	// Cloud Pub/Sub topic to send notifications to. Format is projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#topic GoogleDataLossPreventionDiscoveryConfig#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

