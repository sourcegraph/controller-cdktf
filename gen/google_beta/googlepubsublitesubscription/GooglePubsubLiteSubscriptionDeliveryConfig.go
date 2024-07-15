package googlepubsublitesubscription


type GooglePubsubLiteSubscriptionDeliveryConfig struct {
	// When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_lite_subscription#delivery_requirement GooglePubsubLiteSubscription#delivery_requirement}
	DeliveryRequirement *string `field:"required" json:"deliveryRequirement" yaml:"deliveryRequirement"`
}

