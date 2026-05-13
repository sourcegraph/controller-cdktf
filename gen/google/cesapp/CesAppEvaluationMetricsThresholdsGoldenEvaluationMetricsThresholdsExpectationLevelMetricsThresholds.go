package cesapp


type CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds struct {
	// The success threshold for individual tool invocation parameter correctness. Must be a float between 0 and 1. Default is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#tool_invocation_parameter_correctness_threshold CesApp#tool_invocation_parameter_correctness_threshold}
	ToolInvocationParameterCorrectnessThreshold *float64 `field:"optional" json:"toolInvocationParameterCorrectnessThreshold" yaml:"toolInvocationParameterCorrectnessThreshold"`
}

