package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfig struct {
	// asymmetric_autoscaling_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#asymmetric_autoscaling_options GoogleSpannerInstance#asymmetric_autoscaling_options}
	AsymmetricAutoscalingOptions interface{} `field:"optional" json:"asymmetricAutoscalingOptions" yaml:"asymmetricAutoscalingOptions"`
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#autoscaling_limits GoogleSpannerInstance#autoscaling_limits}
	AutoscalingLimits *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// autoscaling_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_instance#autoscaling_targets GoogleSpannerInstance#autoscaling_targets}
	AutoscalingTargets *GoogleSpannerInstanceAutoscalingConfigAutoscalingTargets `field:"optional" json:"autoscalingTargets" yaml:"autoscalingTargets"`
}

