package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationCsvOptions struct {
	// The value that is used to quote data sections in a CSV file.
	//
	// If your data does not contain quoted sections, set the property value to an empty string. If your data contains quoted newline characters, you must also set the allow_quoted_newlines property to true. The API-side default is ", specified in Terraform escaped as \". Due to limitations with Terraform default values, this value is required to be explicitly set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#quote GoogleBigqueryTable#quote}
	Quote *string `field:"required" json:"quote" yaml:"quote"`
	// Indicates if BigQuery should accept rows that are missing trailing optional columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#allow_jagged_rows GoogleBigqueryTable#allow_jagged_rows}
	AllowJaggedRows interface{} `field:"optional" json:"allowJaggedRows" yaml:"allowJaggedRows"`
	// Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file.
	//
	// The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#allow_quoted_newlines GoogleBigqueryTable#allow_quoted_newlines}
	AllowQuotedNewlines interface{} `field:"optional" json:"allowQuotedNewlines" yaml:"allowQuotedNewlines"`
	// The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#encoding GoogleBigqueryTable#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The separator for fields in a CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#field_delimiter GoogleBigqueryTable#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// The number of rows at the top of a CSV file that BigQuery will skip when reading the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#skip_leading_rows GoogleBigqueryTable#skip_leading_rows}
	SkipLeadingRows *float64 `field:"optional" json:"skipLeadingRows" yaml:"skipLeadingRows"`
}

