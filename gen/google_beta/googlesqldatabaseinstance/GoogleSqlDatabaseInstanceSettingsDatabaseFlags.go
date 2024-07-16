package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsDatabaseFlags struct {
	// Name of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#name GoogleSqlDatabaseInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#value GoogleSqlDatabaseInstance#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

