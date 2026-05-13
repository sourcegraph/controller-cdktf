package googleoracledatabasedbsystem


type GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions struct {
	// Indicates whether to enable data collection for diagnostics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#is_diagnostics_events_enabled GoogleOracleDatabaseDbSystem#is_diagnostics_events_enabled}
	IsDiagnosticsEventsEnabled interface{} `field:"optional" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Indicates whether to enable incident logs and trace collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_db_system#is_incident_logs_enabled GoogleOracleDatabaseDbSystem#is_incident_logs_enabled}
	IsIncidentLogsEnabled interface{} `field:"optional" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

