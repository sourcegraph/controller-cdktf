package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource struct {
	// he name of a Dialogflow virtual agent used for end user side intent detection and suggestion.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#agent GoogleDialogflowConversationProfile#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
	// human_agent_side_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#human_agent_side_config GoogleDialogflowConversationProfile#human_agent_side_config}
	HumanAgentSideConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig `field:"optional" json:"humanAgentSideConfig" yaml:"humanAgentSideConfig"`
}

