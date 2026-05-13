package googledialogflowcxpage


type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings struct {
	// Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#enable_consent_based_redaction GoogleDialogflowCxPage#enable_consent_based_redaction}
	EnableConsentBasedRedaction interface{} `field:"optional" json:"enableConsentBasedRedaction" yaml:"enableConsentBasedRedaction"`
	// Enables DF Interaction logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#enable_interaction_logging GoogleDialogflowCxPage#enable_interaction_logging}
	EnableInteractionLogging interface{} `field:"optional" json:"enableInteractionLogging" yaml:"enableInteractionLogging"`
	// Enables Google Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_page#enable_stackdriver_logging GoogleDialogflowCxPage#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
}

