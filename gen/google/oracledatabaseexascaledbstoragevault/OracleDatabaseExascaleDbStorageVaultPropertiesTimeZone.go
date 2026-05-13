package oracledatabaseexascaledbstoragevault


type OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone struct {
	// IANA Time Zone Database time zone. For example "America/New_York".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#id OracleDatabaseExascaleDbStorageVault#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IANA Time Zone Database version number. For example "2019a".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exascale_db_storage_vault#version OracleDatabaseExascaleDbStorageVault#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

