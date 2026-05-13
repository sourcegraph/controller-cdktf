package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsExpectation struct {
	// agent_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#agent_response GoogleCesEvaluation#agent_response}
	AgentResponse *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse `field:"optional" json:"agentResponse" yaml:"agentResponse"`
	// agent_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#agent_transfer GoogleCesEvaluation#agent_transfer}
	AgentTransfer *GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer `field:"optional" json:"agentTransfer" yaml:"agentTransfer"`
	// mock_tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#mock_tool_response GoogleCesEvaluation#mock_tool_response}
	MockToolResponse *GoogleCesEvaluationGoldenTurnsStepsExpectationMockToolResponse `field:"optional" json:"mockToolResponse" yaml:"mockToolResponse"`
	// A note describing the expectation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#note GoogleCesEvaluation#note}
	Note *string `field:"optional" json:"note" yaml:"note"`
	// tool_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool_call GoogleCesEvaluation#tool_call}
	ToolCall *GoogleCesEvaluationGoldenTurnsStepsExpectationToolCall `field:"optional" json:"toolCall" yaml:"toolCall"`
	// tool_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tool_response GoogleCesEvaluation#tool_response}
	ToolResponse *GoogleCesEvaluationGoldenTurnsStepsExpectationToolResponse `field:"optional" json:"toolResponse" yaml:"toolResponse"`
	// updated_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#updated_variables GoogleCesEvaluation#updated_variables}
	UpdatedVariables *GoogleCesEvaluationGoldenTurnsStepsExpectationUpdatedVariables `field:"optional" json:"updatedVariables" yaml:"updatedVariables"`
}

