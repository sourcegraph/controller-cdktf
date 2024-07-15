package googledialogflowcxtestcase


type GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInput struct {
	// dtmf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#dtmf GoogleDialogflowCxTestCase#dtmf}
	Dtmf *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf `field:"optional" json:"dtmf" yaml:"dtmf"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#event GoogleDialogflowCxTestCase#event}
	Event *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent `field:"optional" json:"event" yaml:"event"`
	// The language of the input.
	//
	// See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes.
	// Note that queries in the same session do not necessarily need to specify the same language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#language_code GoogleDialogflowCxTestCase#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#text GoogleDialogflowCxTestCase#text}
	Text *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText `field:"optional" json:"text" yaml:"text"`
}

