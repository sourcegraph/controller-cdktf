package oracledatabasedbsystem


type OracleDatabaseDbSystemPropertiesDbHomeDatabasePropertiesDbBackupConfig struct {
	// If set to true, enables automatic backups on the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#auto_backup_enabled OracleDatabaseDbSystem#auto_backup_enabled}
	AutoBackupEnabled interface{} `field:"optional" json:"autoBackupEnabled" yaml:"autoBackupEnabled"`
	// Possible values: MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#auto_full_backup_day OracleDatabaseDbSystem#auto_full_backup_day}
	AutoFullBackupDay *string `field:"optional" json:"autoFullBackupDay" yaml:"autoFullBackupDay"`
	// The window in which the full backup should be performed on the database.
	//
	// If no value is provided, the default is anytime.
	// Possible values:
	// SLOT_ONE
	// SLOT_TWO
	// SLOT_THREE
	// SLOT_FOUR
	// SLOT_FIVE
	// SLOT_SIX
	// SLOT_SEVEN
	// SLOT_EIGHT
	// SLOT_NINE
	// SLOT_TEN
	// SLOT_ELEVEN
	// SLOT_TWELVE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#auto_full_backup_window OracleDatabaseDbSystem#auto_full_backup_window}
	AutoFullBackupWindow *string `field:"optional" json:"autoFullBackupWindow" yaml:"autoFullBackupWindow"`
	// The window in which the incremental backup should be performed on the database.
	//
	// If no value is provided, the default is anytime except the auto
	// full backup day.
	// Possible values:
	// SLOT_ONE
	// SLOT_TWO
	// SLOT_THREE
	// SLOT_FOUR
	// SLOT_FIVE
	// SLOT_SIX
	// SLOT_SEVEN
	// SLOT_EIGHT
	// SLOT_NINE
	// SLOT_TEN
	// SLOT_ELEVEN
	// SLOT_TWELVE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#auto_incremental_backup_window OracleDatabaseDbSystem#auto_incremental_backup_window}
	AutoIncrementalBackupWindow *string `field:"optional" json:"autoIncrementalBackupWindow" yaml:"autoIncrementalBackupWindow"`
	// This defines when the backups will be deleted after Database termination. Possible values: DELETE_IMMEDIATELY DELETE_AFTER_RETENTION_PERIOD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#backup_deletion_policy OracleDatabaseDbSystem#backup_deletion_policy}
	BackupDeletionPolicy *string `field:"optional" json:"backupDeletionPolicy" yaml:"backupDeletionPolicy"`
	// backup_destination_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#backup_destination_details OracleDatabaseDbSystem#backup_destination_details}
	BackupDestinationDetails interface{} `field:"optional" json:"backupDestinationDetails" yaml:"backupDestinationDetails"`
	// The number of days an automatic backup is retained before being automatically deleted.
	//
	// This value determines the earliest point in time to
	// which a database can be restored. Min: 1, Max: 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_db_system#retention_period_days OracleDatabaseDbSystem#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
}

