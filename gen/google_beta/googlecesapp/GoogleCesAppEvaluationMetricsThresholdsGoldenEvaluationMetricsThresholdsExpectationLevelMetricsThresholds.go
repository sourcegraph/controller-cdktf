package googlecesapp


type GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds struct {
	// The success threshold for individual tool invocation parameter correctness. Must be a float between 0 and 1. Default is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#tool_invocation_parameter_correctness_threshold GoogleCesApp#tool_invocation_parameter_correctness_threshold}
	ToolInvocationParameterCorrectnessThreshold *float64 `field:"optional" json:"toolInvocationParameterCorrectnessThreshold" yaml:"toolInvocationParameterCorrectnessThreshold"`
}

