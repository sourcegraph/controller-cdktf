package googledatastreamstream


type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning struct {
	// ingestion_time_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#ingestion_time_partition GoogleDatastreamStream#ingestion_time_partition}
	IngestionTimePartition *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition `field:"optional" json:"ingestionTimePartition" yaml:"ingestionTimePartition"`
	// integer_range_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#integer_range_partition GoogleDatastreamStream#integer_range_partition}
	IntegerRangePartition *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition `field:"optional" json:"integerRangePartition" yaml:"integerRangePartition"`
	// If true, queries over the table require a partition filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#require_partition_filter GoogleDatastreamStream#require_partition_filter}
	RequirePartitionFilter interface{} `field:"optional" json:"requirePartitionFilter" yaml:"requirePartitionFilter"`
	// time_unit_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#time_unit_partition GoogleDatastreamStream#time_unit_partition}
	TimeUnitPartition *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition `field:"optional" json:"timeUnitPartition" yaml:"timeUnitPartition"`
}

