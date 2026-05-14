package googlecesevaluation


type GoogleCesEvaluationScenario struct {
	// Rubrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#rubrics GoogleCesEvaluation#rubrics}
	Rubrics *[]*string `field:"required" json:"rubrics" yaml:"rubrics"`
	// scenario_expectations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#scenario_expectations GoogleCesEvaluation#scenario_expectations}
	ScenarioExpectations interface{} `field:"required" json:"scenarioExpectations" yaml:"scenarioExpectations"`
	// The task to evaluate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#task GoogleCesEvaluation#task}
	Task *string `field:"required" json:"task" yaml:"task"`
	// Evaluation expectations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#evaluation_expectations GoogleCesEvaluation#evaluation_expectations}
	EvaluationExpectations *[]*string `field:"optional" json:"evaluationExpectations" yaml:"evaluationExpectations"`
	// Max turns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#max_turns GoogleCesEvaluation#max_turns}
	MaxTurns *float64 `field:"optional" json:"maxTurns" yaml:"maxTurns"`
	// Task completion behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#task_completion_behavior GoogleCesEvaluation#task_completion_behavior}
	TaskCompletionBehavior *string `field:"optional" json:"taskCompletionBehavior" yaml:"taskCompletionBehavior"`
	// user_facts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#user_facts GoogleCesEvaluation#user_facts}
	UserFacts interface{} `field:"optional" json:"userFacts" yaml:"userFacts"`
	// User goal behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#user_goal_behavior GoogleCesEvaluation#user_goal_behavior}
	UserGoalBehavior *string `field:"optional" json:"userGoalBehavior" yaml:"userGoalBehavior"`
	// Variables / Session Parameters as context for the session, keyed by variable names.
	//
	// Members of this struct will override any default values set by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#variable_overrides GoogleCesEvaluation#variable_overrides}
	VariableOverrides *map[string]*string `field:"optional" json:"variableOverrides" yaml:"variableOverrides"`
}

