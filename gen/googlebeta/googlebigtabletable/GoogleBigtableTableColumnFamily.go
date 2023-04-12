package googlebigtabletable


type GoogleBigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_table#family GoogleBigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
}

