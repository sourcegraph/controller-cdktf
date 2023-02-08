package googlepubsublitesubscription


type GooglePubsubLiteSubscriptionDeliveryConfig struct {
	// When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_lite_subscription#delivery_requirement GooglePubsubLiteSubscription#delivery_requirement}
	DeliveryRequirement *string `field:"required" json:"deliveryRequirement" yaml:"deliveryRequirement"`
}

