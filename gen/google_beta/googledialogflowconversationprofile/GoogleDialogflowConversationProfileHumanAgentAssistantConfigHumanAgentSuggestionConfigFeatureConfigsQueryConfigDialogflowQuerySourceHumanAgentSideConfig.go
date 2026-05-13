package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig struct {
	// The name of a dialogflow virtual agent used for intent detection and suggestion triggered by human agent.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#agent GoogleDialogflowConversationProfile#agent}
	Agent *string `field:"optional" json:"agent" yaml:"agent"`
}

