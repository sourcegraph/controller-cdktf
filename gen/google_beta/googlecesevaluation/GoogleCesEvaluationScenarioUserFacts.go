package googlecesevaluation


type GoogleCesEvaluationScenarioUserFacts struct {
	// The name of the user fact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#name GoogleCesEvaluation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the user fact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#value GoogleCesEvaluation#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

