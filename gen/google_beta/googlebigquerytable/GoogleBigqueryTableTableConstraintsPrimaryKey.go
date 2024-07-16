package googlebigquerytable


type GoogleBigqueryTableTableConstraintsPrimaryKey struct {
	// The columns that are composed of the primary key constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#columns GoogleBigqueryTable#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
}

