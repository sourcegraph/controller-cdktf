package googleoracledatabasedbsystem


type GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties struct {
	// The Oracle Database version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#db_version GoogleOracleDatabaseDbSystem#db_version}
	DbVersion *string `field:"required" json:"dbVersion" yaml:"dbVersion"`
	// database_management_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#database_management_config GoogleOracleDatabaseDbSystem#database_management_config}
	DatabaseManagementConfig *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig `field:"optional" json:"databaseManagementConfig" yaml:"databaseManagementConfig"`
	// db_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#db_backup_config GoogleOracleDatabaseDbSystem#db_backup_config}
	DbBackupConfig *GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig `field:"optional" json:"dbBackupConfig" yaml:"dbBackupConfig"`
}

