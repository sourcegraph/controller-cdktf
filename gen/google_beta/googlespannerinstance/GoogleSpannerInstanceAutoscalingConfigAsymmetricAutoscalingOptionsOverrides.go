package googlespannerinstance


type GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides struct {
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance#autoscaling_limits GoogleSpannerInstance#autoscaling_limits}
	AutoscalingLimits *GoogleSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// The target high priority cpu utilization percentage that the autoscaler should be trying to achieve for this replica.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance#autoscaling_target_high_priority_cpu_utilization_percent GoogleSpannerInstance#autoscaling_target_high_priority_cpu_utilization_percent}
	AutoscalingTargetHighPriorityCpuUtilizationPercent *float64 `field:"optional" json:"autoscalingTargetHighPriorityCpuUtilizationPercent" yaml:"autoscalingTargetHighPriorityCpuUtilizationPercent"`
	// The target total cpu utilization percentage that the autoscaler should be trying to achieve for this replica.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance#autoscaling_target_total_cpu_utilization_percent GoogleSpannerInstance#autoscaling_target_total_cpu_utilization_percent}
	AutoscalingTargetTotalCpuUtilizationPercent *float64 `field:"optional" json:"autoscalingTargetTotalCpuUtilizationPercent" yaml:"autoscalingTargetTotalCpuUtilizationPercent"`
	// If true, disables high priority CPU autoscaling for this replica and ignores high_priority_cpu_utilization_percent in the top-level autoscaling configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance#disable_high_priority_cpu_autoscaling GoogleSpannerInstance#disable_high_priority_cpu_autoscaling}
	DisableHighPriorityCpuAutoscaling interface{} `field:"optional" json:"disableHighPriorityCpuAutoscaling" yaml:"disableHighPriorityCpuAutoscaling"`
	// If true, disables total CPU autoscaling for this replica and ignores total_cpu_utilization_percent in the top-level autoscaling configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance#disable_total_cpu_autoscaling GoogleSpannerInstance#disable_total_cpu_autoscaling}
	DisableTotalCpuAutoscaling interface{} `field:"optional" json:"disableTotalCpuAutoscaling" yaml:"disableTotalCpuAutoscaling"`
}

