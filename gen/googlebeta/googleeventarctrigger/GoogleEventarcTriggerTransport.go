package googleeventarctrigger


type GoogleEventarcTriggerTransport struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#pubsub GoogleEventarcTrigger#pubsub}
	Pubsub *GoogleEventarcTriggerTransportPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
}

