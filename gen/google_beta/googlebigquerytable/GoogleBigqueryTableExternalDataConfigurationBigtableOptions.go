package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationBigtableOptions struct {
	// column_family block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#column_family GoogleBigqueryTable#column_family}
	ColumnFamily interface{} `field:"optional" json:"columnFamily" yaml:"columnFamily"`
	// If field is true, then the column families that are not specified in columnFamilies list are not exposed in the table schema.
	//
	// Otherwise, they are read with BYTES type values. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#ignore_unspecified_column_families GoogleBigqueryTable#ignore_unspecified_column_families}
	IgnoreUnspecifiedColumnFamilies interface{} `field:"optional" json:"ignoreUnspecifiedColumnFamilies" yaml:"ignoreUnspecifiedColumnFamilies"`
	// If field is true, then each column family will be read as a single JSON column.
	//
	// Otherwise they are read as a repeated cell structure containing timestamp/value tuples. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#output_column_families_as_json GoogleBigqueryTable#output_column_families_as_json}
	OutputColumnFamiliesAsJson interface{} `field:"optional" json:"outputColumnFamiliesAsJson" yaml:"outputColumnFamiliesAsJson"`
	// If field is true, then the rowkey column families will be read and converted to string.
	//
	// Otherwise they are read with BYTES type values and users need to manually cast them with CAST if necessary. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#read_rowkey_as_string GoogleBigqueryTable#read_rowkey_as_string}
	ReadRowkeyAsString interface{} `field:"optional" json:"readRowkeyAsString" yaml:"readRowkeyAsString"`
}

