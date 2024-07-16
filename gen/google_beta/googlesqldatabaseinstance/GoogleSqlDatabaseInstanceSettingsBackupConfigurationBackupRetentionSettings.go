package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings struct {
	// Number of backups to retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#retained_backups GoogleSqlDatabaseInstance#retained_backups}
	RetainedBackups *float64 `field:"required" json:"retainedBackups" yaml:"retainedBackups"`
	// The unit that 'retainedBackups' represents. Defaults to COUNT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#retention_unit GoogleSqlDatabaseInstance#retention_unit}
	RetentionUnit *string `field:"optional" json:"retentionUnit" yaml:"retentionUnit"`
}

