package oracledatabaseautonomousdatabase


type OracleDatabaseAutonomousDatabaseSourceConfig struct {
	// This field specifies if the replication of automatic backups is enabled when creating a Data Guard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_autonomous_database#automatic_backups_replication_enabled OracleDatabaseAutonomousDatabase#automatic_backups_replication_enabled}
	AutomaticBackupsReplicationEnabled interface{} `field:"optional" json:"automaticBackupsReplicationEnabled" yaml:"automaticBackupsReplicationEnabled"`
	// The name of the primary Autonomous Database that is used to create a Peer Autonomous Database from a source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_autonomous_database#autonomous_database OracleDatabaseAutonomousDatabase#autonomous_database}
	AutonomousDatabase *string `field:"optional" json:"autonomousDatabase" yaml:"autonomousDatabase"`
}

