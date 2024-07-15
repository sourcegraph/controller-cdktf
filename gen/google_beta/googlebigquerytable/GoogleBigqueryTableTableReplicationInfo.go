package googlebigquerytable


type GoogleBigqueryTableTableReplicationInfo struct {
	// The ID of the source dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#source_dataset_id GoogleBigqueryTable#source_dataset_id}
	SourceDatasetId *string `field:"required" json:"sourceDatasetId" yaml:"sourceDatasetId"`
	// The ID of the source project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#source_project_id GoogleBigqueryTable#source_project_id}
	SourceProjectId *string `field:"required" json:"sourceProjectId" yaml:"sourceProjectId"`
	// The ID of the source materialized view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#source_table_id GoogleBigqueryTable#source_table_id}
	SourceTableId *string `field:"required" json:"sourceTableId" yaml:"sourceTableId"`
	// The interval at which the source materialized view is polled for updates. The default is 300000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#replication_interval_ms GoogleBigqueryTable#replication_interval_ms}
	ReplicationIntervalMs *float64 `field:"optional" json:"replicationIntervalMs" yaml:"replicationIntervalMs"`
}

