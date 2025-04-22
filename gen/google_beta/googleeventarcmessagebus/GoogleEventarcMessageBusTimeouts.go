package googleeventarcmessagebus


type GoogleEventarcMessageBusTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_message_bus#create GoogleEventarcMessageBus#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_message_bus#delete GoogleEventarcMessageBus#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_message_bus#update GoogleEventarcMessageBus#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

