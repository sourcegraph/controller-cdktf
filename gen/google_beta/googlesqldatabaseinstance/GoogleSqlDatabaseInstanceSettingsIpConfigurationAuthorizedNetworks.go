package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsIpConfigurationAuthorizedNetworks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#value GoogleSqlDatabaseInstance#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#expiration_time GoogleSqlDatabaseInstance#expiration_time}.
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#name GoogleSqlDatabaseInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

