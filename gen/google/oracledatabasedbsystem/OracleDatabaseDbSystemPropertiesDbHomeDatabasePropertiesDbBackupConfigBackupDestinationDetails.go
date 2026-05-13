package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetails struct {
	// The type of the database backup destination. Possible values: NFS RECOVERY_APPLIANCE OBJECT_STORE LOCAL DBRS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#type OracleDatabaseDbSystem#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

