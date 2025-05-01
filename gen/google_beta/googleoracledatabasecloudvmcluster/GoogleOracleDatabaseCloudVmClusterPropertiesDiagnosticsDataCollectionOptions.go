package googleoracledatabasecloudvmcluster


type GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions struct {
	// Indicates whether diagnostic collection is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_oracle_database_cloud_vm_cluster#diagnostics_events_enabled GoogleOracleDatabaseCloudVmCluster#diagnostics_events_enabled}
	DiagnosticsEventsEnabled interface{} `field:"optional" json:"diagnosticsEventsEnabled" yaml:"diagnosticsEventsEnabled"`
	// Indicates whether health monitoring is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_oracle_database_cloud_vm_cluster#health_monitoring_enabled GoogleOracleDatabaseCloudVmCluster#health_monitoring_enabled}
	HealthMonitoringEnabled interface{} `field:"optional" json:"healthMonitoringEnabled" yaml:"healthMonitoringEnabled"`
	// Indicates whether incident logs and trace collection are enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_oracle_database_cloud_vm_cluster#incident_logs_enabled GoogleOracleDatabaseCloudVmCluster#incident_logs_enabled}
	IncidentLogsEnabled interface{} `field:"optional" json:"incidentLogsEnabled" yaml:"incidentLogsEnabled"`
}

