package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsSteps struct {
	// agent_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#agent_transfer GoogleCesEvaluation#agent_transfer}
	AgentTransfer *GoogleCesEvaluationGoldenTurnsStepsAgentTransfer `field:"optional" json:"agentTransfer" yaml:"agentTransfer"`
	// expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#expectation GoogleCesEvaluation#expectation}
	Expectation *GoogleCesEvaluationGoldenTurnsStepsExpectation `field:"optional" json:"expectation" yaml:"expectation"`
	// user_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#user_input GoogleCesEvaluation#user_input}
	UserInput *GoogleCesEvaluationGoldenTurnsStepsUserInput `field:"optional" json:"userInput" yaml:"userInput"`
}

