package datastreamstream


type DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition struct {
	// The partitioning column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#column DatastreamStream#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// The ending value for range partitioning (exclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#end DatastreamStream#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The interval of each range within the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#interval DatastreamStream#interval}
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The starting value for range partitioning (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#start DatastreamStream#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

