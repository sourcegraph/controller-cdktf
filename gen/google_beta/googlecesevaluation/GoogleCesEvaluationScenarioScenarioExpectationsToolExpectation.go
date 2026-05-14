package googlecesevaluation


type GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation struct {
	// expected_tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#expected_tool_call GoogleCesEvaluation#expected_tool_call}
	ExpectedToolCall *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationExpectedToolCall `field:"optional" json:"expectedToolCall" yaml:"expectedToolCall"`
	// mock_tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#mock_tool_response GoogleCesEvaluation#mock_tool_response}
	MockToolResponse *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectationMockToolResponse `field:"optional" json:"mockToolResponse" yaml:"mockToolResponse"`
}

