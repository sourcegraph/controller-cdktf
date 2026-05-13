package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource struct {
	// he name of a Dialogflow virtual agent used for end user side intent detection and suggestion.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_conversation_profile#agent DialogflowConversationProfile#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
	// human_agent_side_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_conversation_profile#human_agent_side_config DialogflowConversationProfile#human_agent_side_config}
	HumanAgentSideConfig *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig `field:"optional" json:"humanAgentSideConfig" yaml:"humanAgentSideConfig"`
}

