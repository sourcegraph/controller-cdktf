package spannerinstancepartition


type SpannerInstancePartitionAutoscalingConfig struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#autoscaling_limits SpannerInstancePartition#autoscaling_limits}
	AutoscalingLimits *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// autoscaling_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/spanner_instance_partition#autoscaling_targets SpannerInstancePartition#autoscaling_targets}
	AutoscalingTargets *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets `field:"optional" json:"autoscalingTargets" yaml:"autoscalingTargets"`
}

