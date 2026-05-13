package googlecesexample


type GoogleCesExampleMessagesChunksAgentTransfer struct {
	// The agent to which the conversation is being transferred. The agent will handle the conversation from this point forward. Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#target_agent GoogleCesExample#target_agent}
	TargetAgent *string `field:"required" json:"targetAgent" yaml:"targetAgent"`
}

