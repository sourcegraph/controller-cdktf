package backupdrrestoreworkload


type BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#kms_key_name BackupDrRestoreWorkload#kms_key_name}.
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#kms_key_service_account BackupDrRestoreWorkload#kms_key_service_account}.
	KmsKeyServiceAccount *string `field:"optional" json:"kmsKeyServiceAccount" yaml:"kmsKeyServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#raw_key BackupDrRestoreWorkload#raw_key}.
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#rsa_encrypted_key BackupDrRestoreWorkload#rsa_encrypted_key}.
	RsaEncryptedKey *string `field:"optional" json:"rsaEncryptedKey" yaml:"rsaEncryptedKey"`
}

