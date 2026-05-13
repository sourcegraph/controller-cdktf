package oracledatabaseexadbvmcluster


type OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions struct {
	// Indicates whether to enable data collection for diagnostics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#is_diagnostics_events_enabled OracleDatabaseExadbVmCluster#is_diagnostics_events_enabled}
	IsDiagnosticsEventsEnabled interface{} `field:"optional" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Indicates whether to enable health monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#is_health_monitoring_enabled OracleDatabaseExadbVmCluster#is_health_monitoring_enabled}
	IsHealthMonitoringEnabled interface{} `field:"optional" json:"isHealthMonitoringEnabled" yaml:"isHealthMonitoringEnabled"`
	// Indicates whether to enable incident logs and trace collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#is_incident_logs_enabled OracleDatabaseExadbVmCluster#is_incident_logs_enabled}
	IsIncidentLogsEnabled interface{} `field:"optional" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

