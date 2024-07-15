package googledialogflowcxtestcase


type GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput struct {
	// current_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#current_page GoogleDialogflowCxTestCase#current_page}
	CurrentPage *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage `field:"optional" json:"currentPage" yaml:"currentPage"`
	// The session parameters available to the bot at this point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#session_parameters GoogleDialogflowCxTestCase#session_parameters}
	SessionParameters *string `field:"optional" json:"sessionParameters" yaml:"sessionParameters"`
	// text_responses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#text_responses GoogleDialogflowCxTestCase#text_responses}
	TextResponses interface{} `field:"optional" json:"textResponses" yaml:"textResponses"`
	// triggered_intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#triggered_intent GoogleDialogflowCxTestCase#triggered_intent}
	TriggeredIntent *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent `field:"optional" json:"triggeredIntent" yaml:"triggeredIntent"`
}

