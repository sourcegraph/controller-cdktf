package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig struct {
	// Domain name of the Active Directory for SQL Server (e.g., mydomain.com).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_database_instance#domain GoogleSqlDatabaseInstance#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

