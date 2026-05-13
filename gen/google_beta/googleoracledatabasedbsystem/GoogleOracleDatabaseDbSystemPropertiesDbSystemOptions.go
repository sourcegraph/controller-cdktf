package googleoracledatabasedbsystem


type GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions struct {
	// The storage option used in DB system. Possible values: ASM LVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#storage_management GoogleOracleDatabaseDbSystem#storage_management}
	StorageManagement *string `field:"optional" json:"storageManagement" yaml:"storageManagement"`
}

