package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDataCollectionOptions struct {
	// Indicates whether to enable data collection for diagnostics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#is_diagnostics_events_enabled OracleDatabaseDbSystem#is_diagnostics_events_enabled}
	IsDiagnosticsEventsEnabled interface{} `field:"optional" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Indicates whether to enable incident logs and trace collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#is_incident_logs_enabled OracleDatabaseDbSystem#is_incident_logs_enabled}
	IsIncidentLogsEnabled interface{} `field:"optional" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

