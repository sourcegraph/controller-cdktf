package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbSystemOptions struct {
	// The storage option used in DB system. Possible values: ASM LVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#storage_management OracleDatabaseDbSystem#storage_management}
	StorageManagement *string `field:"optional" json:"storageManagement" yaml:"storageManagement"`
}

