package googledatastreamstream


type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition struct {
	// The partitioning column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#column GoogleDatastreamStream#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Partition granularity. Possible values: ["PARTITIONING_TIME_GRANULARITY_UNSPECIFIED", "PARTITIONING_TIME_GRANULARITY_HOUR", "PARTITIONING_TIME_GRANULARITY_DAY", "PARTITIONING_TIME_GRANULARITY_MONTH", "PARTITIONING_TIME_GRANULARITY_YEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#partitioning_time_granularity GoogleDatastreamStream#partitioning_time_granularity}
	PartitioningTimeGranularity *string `field:"optional" json:"partitioningTimeGranularity" yaml:"partitioningTimeGranularity"`
}

