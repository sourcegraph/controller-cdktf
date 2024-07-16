package googledialogflowcxtestcase


type GoogleDialogflowCxTestCaseTestCaseConversationTurns struct {
	// user_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#user_input GoogleDialogflowCxTestCase#user_input}
	UserInput *GoogleDialogflowCxTestCaseTestCaseConversationTurnsUserInput `field:"optional" json:"userInput" yaml:"userInput"`
	// virtual_agent_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#virtual_agent_output GoogleDialogflowCxTestCase#virtual_agent_output}
	VirtualAgentOutput *GoogleDialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput `field:"optional" json:"virtualAgentOutput" yaml:"virtualAgentOutput"`
}

