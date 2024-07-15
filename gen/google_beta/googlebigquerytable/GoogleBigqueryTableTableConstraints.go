package googlebigquerytable


type GoogleBigqueryTableTableConstraints struct {
	// foreign_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#foreign_keys GoogleBigqueryTable#foreign_keys}
	ForeignKeys interface{} `field:"optional" json:"foreignKeys" yaml:"foreignKeys"`
	// primary_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#primary_key GoogleBigqueryTable#primary_key}
	PrimaryKey *GoogleBigqueryTableTableConstraintsPrimaryKey `field:"optional" json:"primaryKey" yaml:"primaryKey"`
}

