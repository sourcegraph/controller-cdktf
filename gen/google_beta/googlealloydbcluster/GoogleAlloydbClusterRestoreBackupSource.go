package googlealloydbcluster


type GoogleAlloydbClusterRestoreBackupSource struct {
	// The name of the backup that this cluster is restored from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#backup_name GoogleAlloydbCluster#backup_name}
	BackupName *string `field:"required" json:"backupName" yaml:"backupName"`
}

