package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#kms_key_name GoogleBackupDrRestoreWorkload#kms_key_name}.
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#kms_key_service_account GoogleBackupDrRestoreWorkload#kms_key_service_account}.
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#raw_key GoogleBackupDrRestoreWorkload#raw_key}.
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#rsa_encrypted_key GoogleBackupDrRestoreWorkload#rsa_encrypted_key}.
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
}

