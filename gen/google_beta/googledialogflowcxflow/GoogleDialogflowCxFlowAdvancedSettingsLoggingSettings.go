package googledialogflowcxflow


type GoogleDialogflowCxFlowAdvancedSettingsLoggingSettings struct {
	// Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#enable_consent_based_redaction GoogleDialogflowCxFlow#enable_consent_based_redaction}
	EnableConsentBasedRedaction interface{} `field:"optional" json:"enableConsentBasedRedaction" yaml:"enableConsentBasedRedaction"`
	// Enables DF Interaction logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#enable_interaction_logging GoogleDialogflowCxFlow#enable_interaction_logging}
	EnableInteractionLogging interface{} `field:"optional" json:"enableInteractionLogging" yaml:"enableInteractionLogging"`
	// Enables Google Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dialogflow_cx_flow#enable_stackdriver_logging GoogleDialogflowCxFlow#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
}

