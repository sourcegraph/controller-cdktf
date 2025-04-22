package spannerinstance


type SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/spanner_instance#autoscaling_limits SpannerInstance#autoscaling_limits}
	AutoscalingLimits *SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits `field:"required" json:"autoscalingLimits" yaml:"autoscalingLimits"`
}

