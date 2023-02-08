package googlepubsublitetopic


type GooglePubsubLiteTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_lite_topic#create GooglePubsubLiteTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_lite_topic#delete GooglePubsubLiteTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_lite_topic#update GooglePubsubLiteTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

