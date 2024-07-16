package googleeventarctrigger


type GoogleEventarcTriggerTransport struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#pubsub GoogleEventarcTrigger#pubsub}
	Pubsub *GoogleEventarcTriggerTransportPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
}

