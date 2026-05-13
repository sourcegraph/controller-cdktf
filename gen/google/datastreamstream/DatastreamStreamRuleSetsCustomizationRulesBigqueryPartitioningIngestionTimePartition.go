package datastreamstream


type DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition struct {
	// Partition granularity. Possible values: ["PARTITIONING_TIME_GRANULARITY_UNSPECIFIED", "PARTITIONING_TIME_GRANULARITY_HOUR", "PARTITIONING_TIME_GRANULARITY_DAY", "PARTITIONING_TIME_GRANULARITY_MONTH", "PARTITIONING_TIME_GRANULARITY_YEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#partitioning_time_granularity DatastreamStream#partitioning_time_granularity}
	PartitioningTimeGranularity *string `field:"optional" json:"partitioningTimeGranularity" yaml:"partitioningTimeGranularity"`
}

