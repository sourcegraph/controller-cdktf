package googledialogflowcxtestcase


type GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInput struct {
	// Whether sentiment analysis is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#enable_sentiment_analysis GoogleDialogflowCxTestCase#enable_sentiment_analysis}
	EnableSentimentAnalysis interface{} `field:"optional" json:"enableSentimentAnalysis" yaml:"enableSentimentAnalysis"`
	// Parameters that need to be injected into the conversation during intent detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#injected_parameters GoogleDialogflowCxTestCase#injected_parameters}
	InjectedParameters *string `field:"optional" json:"injectedParameters" yaml:"injectedParameters"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#input GoogleDialogflowCxTestCase#input}
	Input *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput `field:"optional" json:"input" yaml:"input"`
	// If webhooks should be allowed to trigger in response to the user utterance.
	//
	// Often if parameters are injected, webhooks should not be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#is_webhook_enabled GoogleDialogflowCxTestCase#is_webhook_enabled}
	IsWebhookEnabled interface{} `field:"optional" json:"isWebhookEnabled" yaml:"isWebhookEnabled"`
}

