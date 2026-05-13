package googlebackupdrbackupvault


type GoogleBackupDrBackupVaultEncryptionConfig struct {
	// The Resource name of the Cloud KMS key to be used to encrypt new backups.
	//
	// The key must be in the same location as the backup vault. The key must be a Cloud KMS CryptoKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_backup_vault#kms_key_name GoogleBackupDrBackupVault#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

