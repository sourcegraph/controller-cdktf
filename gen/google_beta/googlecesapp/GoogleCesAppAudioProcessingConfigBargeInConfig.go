package googlecesapp


type GoogleCesAppAudioProcessingConfigBargeInConfig struct {
	// If enabled, the agent will adapt its next response based on the assumption that the user hasn't heard the full preceding agent message.
	//
	// This should not be used in scenarios where agent responses are displayed
	// visually.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#barge_in_awareness GoogleCesApp#barge_in_awareness}
	BargeInAwareness interface{} `field:"optional" json:"bargeInAwareness" yaml:"bargeInAwareness"`
}

