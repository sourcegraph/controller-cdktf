package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings struct {
	// If set to true, the last message from virtual agent (hand off message) and the message before it (trigger message of hand off) are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#drop_handoff_messages GoogleDialogflowConversationProfile#drop_handoff_messages}
	DropHandoffMessages interface{} `field:"optional" json:"dropHandoffMessages" yaml:"dropHandoffMessages"`
	// If set to true, all messages from ivr stage are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#drop_ivr_messages GoogleDialogflowConversationProfile#drop_ivr_messages}
	DropIvrMessages interface{} `field:"optional" json:"dropIvrMessages" yaml:"dropIvrMessages"`
	// If set to true, all messages from virtual agent are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#drop_virtual_agent_messages GoogleDialogflowConversationProfile#drop_virtual_agent_messages}
	DropVirtualAgentMessages interface{} `field:"optional" json:"dropVirtualAgentMessages" yaml:"dropVirtualAgentMessages"`
}

