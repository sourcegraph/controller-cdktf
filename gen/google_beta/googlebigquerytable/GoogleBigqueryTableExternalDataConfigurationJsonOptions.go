package googlebigquerytable


type GoogleBigqueryTableExternalDataConfigurationJsonOptions struct {
	// The character encoding of the data.
	//
	// The supported values are UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, and UTF-32LE. The default value is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#encoding GoogleBigqueryTable#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
}

