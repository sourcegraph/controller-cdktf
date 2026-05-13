package alloydbcluster


type AlloydbClusterRestoreBackupdrPitrSource struct {
	// The name of the BackupDR data source that this cluster is restore from.
	//
	// It must be of the format "projects/[PROJECT]/locations/[LOCATION]/backupVaults/[VAULT_ID]/dataSources/[DATASOURCE_ID]"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_cluster#data_source AlloydbCluster#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The point in time that this cluster is restored to, in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_cluster#point_in_time AlloydbCluster#point_in_time}
	PointInTime *string `field:"required" json:"pointInTime" yaml:"pointInTime"`
}

