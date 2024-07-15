package googlebigtabletable


type GoogleBigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_table#family GoogleBigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
}

