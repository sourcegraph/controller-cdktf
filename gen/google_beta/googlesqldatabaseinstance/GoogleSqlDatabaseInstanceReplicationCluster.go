package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceReplicationCluster struct {
	// If the instance is a primary instance, then this field identifies the disaster recovery (DR) replica.
	//
	// The standard format of this field is "your-project:your-instance". You can also set this field to "your-instance", but cloud SQL backend will convert it to the aforementioned standard format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_sql_database_instance#failover_dr_replica_name GoogleSqlDatabaseInstance#failover_dr_replica_name}
	FailoverDrReplicaName *string `field:"optional" json:"failoverDrReplicaName" yaml:"failoverDrReplicaName"`
}

