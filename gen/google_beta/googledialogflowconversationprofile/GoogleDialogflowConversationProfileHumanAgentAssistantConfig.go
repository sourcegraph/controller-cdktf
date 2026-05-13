package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfig struct {
	// end_user_suggestion_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#end_user_suggestion_config GoogleDialogflowConversationProfile#end_user_suggestion_config}
	EndUserSuggestionConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig `field:"optional" json:"endUserSuggestionConfig" yaml:"endUserSuggestionConfig"`
	// human_agent_suggestion_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#human_agent_suggestion_config GoogleDialogflowConversationProfile#human_agent_suggestion_config}
	HumanAgentSuggestionConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig `field:"optional" json:"humanAgentSuggestionConfig" yaml:"humanAgentSuggestionConfig"`
	// message_analysis_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#message_analysis_config GoogleDialogflowConversationProfile#message_analysis_config}
	MessageAnalysisConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfig `field:"optional" json:"messageAnalysisConfig" yaml:"messageAnalysisConfig"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#notification_config GoogleDialogflowConversationProfile#notification_config}
	NotificationConfig *GoogleDialogflowConversationProfileHumanAgentAssistantConfigNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
}

