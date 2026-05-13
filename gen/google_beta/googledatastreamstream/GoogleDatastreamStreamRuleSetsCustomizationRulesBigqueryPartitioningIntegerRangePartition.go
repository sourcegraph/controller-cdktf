package googledatastreamstream


type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition struct {
	// The partitioning column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#column GoogleDatastreamStream#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// The ending value for range partitioning (exclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#end GoogleDatastreamStream#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The interval of each range within the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#interval GoogleDatastreamStream#interval}
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The starting value for range partitioning (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#start GoogleDatastreamStream#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

