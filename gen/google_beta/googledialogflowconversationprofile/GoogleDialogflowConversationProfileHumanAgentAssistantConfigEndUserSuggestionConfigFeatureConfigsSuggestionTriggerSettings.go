package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings struct {
	// Do not trigger if last utterance is small talk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#no_small_talk GoogleDialogflowConversationProfile#no_small_talk}
	NoSmallTalk interface{} `field:"optional" json:"noSmallTalk" yaml:"noSmallTalk"`
	// Only trigger suggestion if participant role of last utterance is END_USER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#only_end_user GoogleDialogflowConversationProfile#only_end_user}
	OnlyEndUser interface{} `field:"optional" json:"onlyEndUser" yaml:"onlyEndUser"`
}

