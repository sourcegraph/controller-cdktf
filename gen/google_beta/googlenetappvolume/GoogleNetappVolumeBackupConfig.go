package googlenetappvolume


type GoogleNetappVolumeBackupConfig struct {
	// Specify a single backup policy ID for scheduled backups. Format: 'projects/{{projectId}}/locations/{{location}}/backupPolicies/{{backupPolicyName}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#backup_policies GoogleNetappVolume#backup_policies}
	BackupPolicies *[]*string `field:"optional" json:"backupPolicies" yaml:"backupPolicies"`
	// ID of the backup vault to use. A backup vault is reqired to create manual or scheduled backups. Format: 'projects/{{projectId}}/locations/{{location}}/backupVaults/{{backupVaultName}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#backup_vault GoogleNetappVolume#backup_vault}
	BackupVault *string `field:"optional" json:"backupVault" yaml:"backupVault"`
	// When set to true, scheduled backup is enabled on the volume. Omit if no backup_policy is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#scheduled_backup_enabled GoogleNetappVolume#scheduled_backup_enabled}
	ScheduledBackupEnabled interface{} `field:"optional" json:"scheduledBackupEnabled" yaml:"scheduledBackupEnabled"`
}

