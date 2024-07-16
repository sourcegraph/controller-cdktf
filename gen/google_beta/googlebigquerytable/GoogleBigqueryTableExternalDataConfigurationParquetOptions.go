package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationParquetOptions struct {
	// Indicates whether to use schema inference specifically for Parquet LIST logical type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#enable_list_inference GoogleBigqueryTable#enable_list_inference}
	EnableListInference interface{} `field:"optional" json:"enableListInference" yaml:"enableListInference"`
	// Indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#enum_as_string GoogleBigqueryTable#enum_as_string}
	EnumAsString interface{} `field:"optional" json:"enumAsString" yaml:"enumAsString"`
}

