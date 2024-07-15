package googlebigquerytable


type GoogleBigqueryTableTableConstraintsForeignKeys struct {
	// column_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#column_references GoogleBigqueryTable#column_references}
	ColumnReferences *GoogleBigqueryTableTableConstraintsForeignKeysColumnReferences `field:"required" json:"columnReferences" yaml:"columnReferences"`
	// referenced_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#referenced_table GoogleBigqueryTable#referenced_table}
	ReferencedTable *GoogleBigqueryTableTableConstraintsForeignKeysReferencedTable `field:"required" json:"referencedTable" yaml:"referencedTable"`
	// Set only if the foreign key constraint is named.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#name GoogleBigqueryTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

