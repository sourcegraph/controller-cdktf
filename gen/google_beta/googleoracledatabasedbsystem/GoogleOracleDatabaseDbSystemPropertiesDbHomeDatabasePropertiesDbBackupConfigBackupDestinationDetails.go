package googleoracledatabasedbsystem


type GoogleOracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfigBackupDestinationDetails struct {
	// The type of the database backup destination. Possible values: NFS RECOVERY_APPLIANCE OBJECT_STORE LOCAL DBRS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#type GoogleOracleDatabaseDbSystem#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

