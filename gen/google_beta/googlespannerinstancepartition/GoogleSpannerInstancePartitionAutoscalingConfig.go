package googlespannerinstancepartition


type GoogleSpannerInstancePartitionAutoscalingConfig struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#autoscaling_limits GoogleSpannerInstancePartition#autoscaling_limits}
	AutoscalingLimits *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// autoscaling_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#autoscaling_targets GoogleSpannerInstancePartition#autoscaling_targets}
	AutoscalingTargets *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets `field:"optional" json:"autoscalingTargets" yaml:"autoscalingTargets"`
}

