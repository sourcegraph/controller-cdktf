package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfig struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#autoscaling_limits GoogleSpannerInstance#autoscaling_limits}
	AutoscalingLimits *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// autoscaling_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#autoscaling_targets GoogleSpannerInstance#autoscaling_targets}
	AutoscalingTargets *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets `field:"optional" json:"autoscalingTargets" yaml:"autoscalingTargets"`
}

