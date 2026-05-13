package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessages struct {
	// The channel which the response is associated with.
	//
	// Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#channel GoogleDialogflowCxPage#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// conversation_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#conversation_success GoogleDialogflowCxPage#conversation_success}
	ConversationSuccess *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess `field:"optional" json:"conversationSuccess" yaml:"conversationSuccess"`
	// knowledge_info_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#knowledge_info_card GoogleDialogflowCxPage#knowledge_info_card}
	KnowledgeInfoCard *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard `field:"optional" json:"knowledgeInfoCard" yaml:"knowledgeInfoCard"`
	// live_agent_handoff block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#live_agent_handoff GoogleDialogflowCxPage#live_agent_handoff}
	LiveAgentHandoff *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff `field:"optional" json:"liveAgentHandoff" yaml:"liveAgentHandoff"`
	// output_audio_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#output_audio_text GoogleDialogflowCxPage#output_audio_text}
	OutputAudioText *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText `field:"optional" json:"outputAudioText" yaml:"outputAudioText"`
	// Returns a response containing a custom, platform-specific payload.
	//
	// This field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#payload GoogleDialogflowCxPage#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// play_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#play_audio GoogleDialogflowCxPage#play_audio}
	PlayAudio *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio `field:"optional" json:"playAudio" yaml:"playAudio"`
	// telephony_transfer_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#telephony_transfer_call GoogleDialogflowCxPage#telephony_transfer_call}
	TelephonyTransferCall *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall `field:"optional" json:"telephonyTransferCall" yaml:"telephonyTransferCall"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#text GoogleDialogflowCxPage#text}
	Text *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText `field:"optional" json:"text" yaml:"text"`
}

