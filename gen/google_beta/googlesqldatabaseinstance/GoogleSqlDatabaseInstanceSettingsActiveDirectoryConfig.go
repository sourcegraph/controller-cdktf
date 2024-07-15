package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig struct {
	// Domain name of the Active Directory for SQL Server (e.g., mydomain.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#domain GoogleSqlDatabaseInstance#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

