package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsFinalBackupConfig struct {
	// When this parameter is set to true, the final backup is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#enabled GoogleSqlDatabaseInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The number of days to retain the final backup after the instance deletion.
	//
	// The valid range is between 1 and 365. For instances managed by BackupDR, the valid range is between 1 day and 99 years. The final backup will be purged at (time_of_instance_deletion + retention_days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_sql_database_instance#retention_days GoogleSqlDatabaseInstance#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

