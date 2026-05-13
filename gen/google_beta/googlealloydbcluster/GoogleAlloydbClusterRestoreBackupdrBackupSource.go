package googlealloydbcluster


type GoogleAlloydbClusterRestoreBackupdrBackupSource struct {
	// The name of the BackupDR backup that this cluster is restored from. It must be of the format "projects/[PROJECT]/locations/[LOCATION]/backupVaults/[VAULT_ID]/dataSources/[DATASOURCE_ID]/backups/[BACKUP_ID]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_alloydb_cluster#backup GoogleAlloydbCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

