package googlecesevaluation


type GoogleCesEvaluationGoldenTurnsStepsExpectationAgentTransfer struct {
	// The display name of the target agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#display_name GoogleCesEvaluation#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The resource name of the target agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#target_agent GoogleCesEvaluation#target_agent}
	TargetAgent *string `field:"optional" json:"targetAgent" yaml:"targetAgent"`
}

