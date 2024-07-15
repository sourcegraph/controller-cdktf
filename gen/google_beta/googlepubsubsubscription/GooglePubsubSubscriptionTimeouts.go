package googlepubsubsubscription


type GooglePubsubSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription#create GooglePubsubSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription#delete GooglePubsubSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription#update GooglePubsubSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

