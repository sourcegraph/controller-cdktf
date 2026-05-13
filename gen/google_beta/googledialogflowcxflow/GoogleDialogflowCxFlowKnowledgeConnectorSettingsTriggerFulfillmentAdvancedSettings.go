package googledialogflowcxflow


type GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings struct {
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#dtmf_settings GoogleDialogflowCxFlow#dtmf_settings}
	DtmfSettings *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
	// logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#logging_settings GoogleDialogflowCxFlow#logging_settings}
	LoggingSettings *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings `field:"optional" json:"loggingSettings" yaml:"loggingSettings"`
	// speech_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_flow#speech_settings GoogleDialogflowCxFlow#speech_settings}
	SpeechSettings *GoogleDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsSpeechSettings `field:"optional" json:"speechSettings" yaml:"speechSettings"`
}

