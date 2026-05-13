package cesapp


type CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds struct {
	// expectation_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#expectation_level_metrics_thresholds CesApp#expectation_level_metrics_thresholds}
	ExpectationLevelMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds `field:"optional" json:"expectationLevelMetricsThresholds" yaml:"expectationLevelMetricsThresholds"`
	// turn_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#turn_level_metrics_thresholds CesApp#turn_level_metrics_thresholds}
	TurnLevelMetricsThresholds *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds `field:"optional" json:"turnLevelMetricsThresholds" yaml:"turnLevelMetricsThresholds"`
}

