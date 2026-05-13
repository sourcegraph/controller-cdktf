package alloydbcluster


type AlloydbClusterRestoreBackupdrBackupSource struct {
	// The name of the BackupDR backup that this cluster is restored from. It must be of the format "projects/[PROJECT]/locations/[LOCATION]/backupVaults/[VAULT_ID]/dataSources/[DATASOURCE_ID]/backups/[BACKUP_ID]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_cluster#backup AlloydbCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

