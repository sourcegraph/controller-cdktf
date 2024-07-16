package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumn struct {
	// The encoding of the values when the type is not STRING.
	//
	// Acceptable encoding values are: TEXT - indicates values are alphanumeric text strings. BINARY - indicates values are encoded using HBase Bytes.toBytes family of functions. 'encoding' can also be set at the column family level. However, the setting at this level takes precedence if 'encoding' is set at both levels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#encoding GoogleBigqueryTable#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// If the qualifier is not a valid BigQuery field identifier i.e. does not match [a-zA-Z][a-zA-Z0-9_]*, a valid identifier must be provided as the column field name and is used as field name in queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#field_name GoogleBigqueryTable#field_name}
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// If this is set, only the latest version of value in this column are exposed.
	//
	// 'onlyReadLatest' can also be set at the column family level. However, the setting at this level takes precedence if 'onlyReadLatest' is set at both levels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#only_read_latest GoogleBigqueryTable#only_read_latest}
	OnlyReadLatest interface{} `field:"optional" json:"onlyReadLatest" yaml:"onlyReadLatest"`
	// Qualifier of the column.
	//
	// Columns in the parent column family that has this exact qualifier are exposed as . field. If the qualifier is valid UTF-8 string, it can be specified in the qualifierString field. Otherwise, a base-64 encoded value must be set to qualifierEncoded. The column field name is the same as the column qualifier. However, if the qualifier is not a valid BigQuery field identifier i.e. does not match [a-zA-Z][a-zA-Z0-9_]*, a valid identifier must be provided as fieldName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#qualifier_encoded GoogleBigqueryTable#qualifier_encoded}
	QualifierEncoded *string `field:"optional" json:"qualifierEncoded" yaml:"qualifierEncoded"`
	// Qualifier string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#qualifier_string GoogleBigqueryTable#qualifier_string}
	QualifierString *string `field:"optional" json:"qualifierString" yaml:"qualifierString"`
	// The type to convert the value in cells of this column.
	//
	// The values are expected to be encoded using HBase Bytes.toBytes function when using the BINARY encoding value. Following BigQuery types are allowed (case-sensitive): "BYTES", "STRING", "INTEGER", "FLOAT", "BOOLEAN", "JSON", Default type is "BYTES". 'type' can also be set at the column family level. However, the setting at this level takes precedence if 'type' is set at both levels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#type GoogleBigqueryTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

