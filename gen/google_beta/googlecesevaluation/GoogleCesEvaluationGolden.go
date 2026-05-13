package googlecesevaluation


type GoogleCesEvaluationGolden struct {
	// turns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#turns GoogleCesEvaluation#turns}
	Turns interface{} `field:"required" json:"turns" yaml:"turns"`
	// The evaluation expectations to evaluate the replayed conversation against. Format: projects/{project}/locations/{location}/apps/{app}/evaluationExpectations/{evaluationExpectation}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#evaluation_expectations GoogleCesEvaluation#evaluation_expectations}
	EvaluationExpectations *[]*string `field:"optional" json:"evaluationExpectations" yaml:"evaluationExpectations"`
}

