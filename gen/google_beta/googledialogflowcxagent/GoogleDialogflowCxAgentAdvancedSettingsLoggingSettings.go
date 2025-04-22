package googledialogflowcxagent


type GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings struct {
	// Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dialogflow_cx_agent#enable_consent_based_redaction GoogleDialogflowCxAgent#enable_consent_based_redaction}
	EnableConsentBasedRedaction interface{} `field:"optional" json:"enableConsentBasedRedaction" yaml:"enableConsentBasedRedaction"`
	// Enables DF Interaction logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dialogflow_cx_agent#enable_interaction_logging GoogleDialogflowCxAgent#enable_interaction_logging}
	EnableInteractionLogging interface{} `field:"optional" json:"enableInteractionLogging" yaml:"enableInteractionLogging"`
	// Enables Google Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dialogflow_cx_agent#enable_stackdriver_logging GoogleDialogflowCxAgent#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
}

