package googlebigquerytable


type GoogleBigqueryTableTableConstraintsForeignKeysColumnReferences struct {
	// The column in the primary key that are referenced by the referencingColumn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#referenced_column GoogleBigqueryTable#referenced_column}
	ReferencedColumn *string `field:"required" json:"referencedColumn" yaml:"referencedColumn"`
	// The column that composes the foreign key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#referencing_column GoogleBigqueryTable#referencing_column}
	ReferencingColumn *string `field:"required" json:"referencingColumn" yaml:"referencingColumn"`
}

