package googlebigquerytable


type GoogleBigqueryTableTimePartitioning struct {
	// The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#type GoogleBigqueryTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Number of milliseconds for which to keep the storage for a partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#expiration_ms GoogleBigqueryTable#expiration_ms}
	ExpirationMs *float64 `field:"optional" json:"expirationMs" yaml:"expirationMs"`
	// The field used to determine how to create a time-based partition.
	//
	// If time-based partitioning is enabled without this value, the table is partitioned based on the load time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#field GoogleBigqueryTable#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#require_partition_filter GoogleBigqueryTable#require_partition_filter}
	RequirePartitionFilter interface{} `field:"optional" json:"requirePartitionFilter" yaml:"requirePartitionFilter"`
}

