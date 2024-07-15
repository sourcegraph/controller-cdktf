package googlenetappvolume


type GoogleNetappVolumeRestoreParameters struct {
	// Full name of the snapshot to use for creating this volume. 'source_snapshot' and 'source_backup' cannot be used simultaneously. Format: 'projects/{{project}}/locations/{{location}}/backupVaults/{{backupVaultId}}/backups/{{backup}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#source_backup GoogleNetappVolume#source_backup}
	SourceBackup *string `field:"optional" json:"sourceBackup" yaml:"sourceBackup"`
	// Full name of the snapshot to use for creating this volume. 'source_snapshot' and 'source_backup' cannot be used simultaneously. Format: 'projects/{{project}}/locations/{{location}}/volumes/{{volume}}/snapshots/{{snapshot}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#source_snapshot GoogleNetappVolume#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
}

