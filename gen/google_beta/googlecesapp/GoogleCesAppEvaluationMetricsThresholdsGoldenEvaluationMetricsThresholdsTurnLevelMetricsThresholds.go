package googlecesapp


type GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds struct {
	// The success threshold for overall tool invocation correctness. Must be a float between 0 and 1. Default is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#overall_tool_invocation_correctness_threshold GoogleCesApp#overall_tool_invocation_correctness_threshold}
	OverallToolInvocationCorrectnessThreshold *float64 `field:"optional" json:"overallToolInvocationCorrectnessThreshold" yaml:"overallToolInvocationCorrectnessThreshold"`
	// The success threshold for semantic similarity. Must be an integer between 0 and 4. Default is >= 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#semantic_similarity_success_threshold GoogleCesApp#semantic_similarity_success_threshold}
	SemanticSimilaritySuccessThreshold *float64 `field:"optional" json:"semanticSimilaritySuccessThreshold" yaml:"semanticSimilaritySuccessThreshold"`
}

