package googlebigquerytable


type GoogleBigqueryTableRangePartitioningRange struct {
	// End of the range partitioning, exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#end GoogleBigqueryTable#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The width of each range within the partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#interval GoogleBigqueryTable#interval}
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Start of the range partitioning, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#start GoogleBigqueryTable#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

