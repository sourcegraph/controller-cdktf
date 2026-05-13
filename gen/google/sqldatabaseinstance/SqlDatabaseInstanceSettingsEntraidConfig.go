package sqldatabaseinstance


type SqlDatabaseInstanceSettingsEntraidConfig struct {
	// The application ID for the Entra ID configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#application_id SqlDatabaseInstance#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The tenant ID for the Entra ID configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#tenant_id SqlDatabaseInstance#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

