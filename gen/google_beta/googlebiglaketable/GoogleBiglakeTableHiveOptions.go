package googlebiglaketable


type GoogleBiglakeTableHiveOptions struct {
	// Stores user supplied Hive table parameters.
	//
	// An object containing a
	// list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#parameters GoogleBiglakeTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#storage_descriptor GoogleBiglakeTable#storage_descriptor}
	StorageDescriptor *GoogleBiglakeTableHiveOptionsStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#table_type GoogleBiglakeTable#table_type}
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
}

