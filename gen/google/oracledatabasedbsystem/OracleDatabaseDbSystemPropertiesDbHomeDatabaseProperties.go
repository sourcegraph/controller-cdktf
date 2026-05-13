package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbHomeDatabaseProperties struct {
	// The Oracle Database version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_version OracleDatabaseDbSystem#db_version}
	DbVersion *string `field:"required" json:"dbVersion" yaml:"dbVersion"`
	// database_management_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#database_management_config OracleDatabaseDbSystem#database_management_config}
	DatabaseManagementConfig *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDatabaseManagementConfig `field:"optional" json:"databaseManagementConfig" yaml:"databaseManagementConfig"`
	// db_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#db_backup_config OracleDatabaseDbSystem#db_backup_config}
	DbBackupConfig *OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig `field:"optional" json:"dbBackupConfig" yaml:"dbBackupConfig"`
}

