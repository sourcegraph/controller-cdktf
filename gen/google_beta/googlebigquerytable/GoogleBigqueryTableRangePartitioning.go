package googlebigquerytable


type GoogleBigqueryTableRangePartitioning struct {
	// The field used to determine how to create a range-based partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#field GoogleBigqueryTable#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#range GoogleBigqueryTable#range}
	Range *GoogleBigqueryTableRangePartitioningRange `field:"required" json:"range" yaml:"range"`
}

