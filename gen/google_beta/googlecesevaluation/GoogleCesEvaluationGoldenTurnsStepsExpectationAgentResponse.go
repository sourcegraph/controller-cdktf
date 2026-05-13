package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsExpectationAgentResponse struct {
	// chunks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#chunks GoogleCesEvaluation#chunks}
	Chunks interface{} `field:"optional" json:"chunks" yaml:"chunks"`
	// The role within the conversation, e.g., user, agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#role GoogleCesEvaluation#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

