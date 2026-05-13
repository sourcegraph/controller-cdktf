package googlespannerinstancepartition


type GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets struct {
	// Specifies the target high priority cpu utilization percentage that the autoscaler should be trying to achieve for the instance partition.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#high_priority_cpu_utilization_percent GoogleSpannerInstancePartition#high_priority_cpu_utilization_percent}
	HighPriorityCpuUtilizationPercent *float64 `field:"optional" json:"highPriorityCpuUtilizationPercent" yaml:"highPriorityCpuUtilizationPercent"`
	// Specifies the target storage utilization percentage that the autoscaler should be trying to achieve for the instance partition.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#storage_utilization_percent GoogleSpannerInstancePartition#storage_utilization_percent}
	StorageUtilizationPercent *float64 `field:"optional" json:"storageUtilizationPercent" yaml:"storageUtilizationPercent"`
	// Specifies the target total cpu utilization percentage that the autoscaler should be trying to achieve for the instance partition.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization). The valid range is [10, 90] inclusive.
	// If not specified or set to 0, the autoscaler will skip scaling based on total cpu utilization.
	// The value should be higher than high_priority_cpu_utilization_percent if present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_spanner_instance_partition#total_cpu_utilization_percent GoogleSpannerInstancePartition#total_cpu_utilization_percent}
	TotalCpuUtilizationPercent *float64 `field:"optional" json:"totalCpuUtilizationPercent" yaml:"totalCpuUtilizationPercent"`
}

