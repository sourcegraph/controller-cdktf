package googlespannerbackupschedule


type GoogleSpannerBackupScheduleEncryptionConfig struct {
	// The encryption type of backups created by the backup schedule.
	//
	// Possible values are USE_DATABASE_ENCRYPTION, GOOGLE_DEFAULT_ENCRYPTION, or CUSTOMER_MANAGED_ENCRYPTION.
	// If you use CUSTOMER_MANAGED_ENCRYPTION, you must specify a kmsKeyName.
	// If your backup type is incremental-backup, the encryption type must be GOOGLE_DEFAULT_ENCRYPTION. Possible values: ["USE_DATABASE_ENCRYPTION", "GOOGLE_DEFAULT_ENCRYPTION", "CUSTOMER_MANAGED_ENCRYPTION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_spanner_backup_schedule#encryption_type GoogleSpannerBackupSchedule#encryption_type}
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// The resource name of the Cloud KMS key to use for encryption. Format: 'projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_spanner_backup_schedule#kms_key_name GoogleSpannerBackupSchedule#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

