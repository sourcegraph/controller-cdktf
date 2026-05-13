package datastreamstream


type DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning struct {
	// ingestion_time_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#ingestion_time_partition DatastreamStream#ingestion_time_partition}
	IngestionTimePartition *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition `field:"optional" json:"ingestionTimePartition" yaml:"ingestionTimePartition"`
	// integer_range_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#integer_range_partition DatastreamStream#integer_range_partition}
	IntegerRangePartition *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition `field:"optional" json:"integerRangePartition" yaml:"integerRangePartition"`
	// If true, queries over the table require a partition filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#require_partition_filter DatastreamStream#require_partition_filter}
	RequirePartitionFilter interface{} `field:"optional" json:"requirePartitionFilter" yaml:"requirePartitionFilter"`
	// time_unit_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#time_unit_partition DatastreamStream#time_unit_partition}
	TimeUnitPartition *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition `field:"optional" json:"timeUnitPartition" yaml:"timeUnitPartition"`
}

