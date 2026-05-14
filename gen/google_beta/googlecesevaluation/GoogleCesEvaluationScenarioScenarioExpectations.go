package googlecesevaluation


type GoogleCesEvaluationScenarioScenarioExpectations struct {
	// agent_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#agent_response GoogleCesEvaluation#agent_response}
	AgentResponse *GoogleCesEvaluationScenarioScenarioExpectationsAgentResponse `field:"optional" json:"agentResponse" yaml:"agentResponse"`
	// tool_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#tool_expectation GoogleCesEvaluation#tool_expectation}
	ToolExpectation *GoogleCesEvaluationScenarioScenarioExpectationsToolExpectation `field:"optional" json:"toolExpectation" yaml:"toolExpectation"`
}

